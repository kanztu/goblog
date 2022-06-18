package controller

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/pkg/ginrunner"
	"github.com/kanztu/goblog/pkg/model"
	"github.com/kanztu/goblog/pkg/server_context"
)

type BlogController struct {
}

func NewBlogController() *BlogController {
	return &BlogController{}
}

// GetBlog
// @Summary      Get Blog
// @Description  GetBlog
// @Tags         Blogs
// @Produce      json
// @Success      200  {object}  []BlogRsp
// @Router       /blogs [get]
func (c *BlogController) GetBlog(ctx *gin.Context) {
	var rsp []BlogRsp
	var blogtags []model.BlogTag

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()

	session.Table("blog").Join("INNER", "tag", "blog.tag_id = tag.tag_id").Find(&blogtags)
	server_context.SrvCtx.Logger.Debugf("%d blogs found in db", len(blogtags))
	for _, v := range blogtags {
		var r BlogRsp
		r.Id = v.Id
		r.Title = v.Title
		r.Auther = v.Auther
		r.Tag = v.TagName
		r.TagId = v.Tag.TagId
		r.Icon = v.Icon
		r.Description = v.Description
		r.CreatedAt = v.CreatedAt
		rsp = append(rsp, r)
	}

	ginrunner.ResponseJSON(ctx, nil, rsp)
}

// GetBlog
// @Summary      Get Blog with tag
// @Description  GetBlog tag
// @Tags         Blogs
// @Accept       json
// @Param        GetBlogReq   body      GetBlogReq  true  "GetBlogReq"
// @Produce      json
// @Success      200  {object}  []BlogRsp
// @Router       /blogs/tag [get]
func (c *BlogController) GetBlogWithTag(ctx *gin.Context) {
	var req GetBlogReq
	var rsp []BlogRsp
	var blogtags []model.BlogTag

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()

	if err := ctx.BindJSON(&req); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	sql := session.Table("blog").Join("INNER", "tag", "blog.tag_id = tag.tag_id")
	if req.TagId >= 0 {
		sql = sql.Where("blog.tag_id = ?", req.TagId)
	}
	sql.Find(&blogtags)
	for _, v := range blogtags {
		var r BlogRsp
		r.Id = v.Id
		r.Title = v.Title
		r.Auther = v.Auther
		r.Tag = v.TagName
		r.TagId = v.Tag.TagId
		r.Icon = v.Icon
		r.Description = v.Description
		r.CreatedAt = v.CreatedAt
		rsp = append(rsp, r)
	}

	ginrunner.ResponseJSON(ctx, nil, rsp)
}

// Create Blog
// @Summary      Create Blog
// @Description  Create Blog
// @Tags         Blogs
// @Accept       json
// @Param        CreateBlogReq   body      CreateBlogReq  true  "CreateBlogReq"
// @Produce      json
// @Success      200  {object}  BlogRsp
// @Router       /admin/blogs [post]
func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var err error
	var req CreateBlogReq
	var t model.Tag
	var b model.Blog
	var content model.Content

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	if err = ctx.BindJSON(&req); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	if req.TagName == "" && req.TagId <= 0 {
		ginrunner.ResponseJSON(ctx, errors.New("no tag provided"), nil)
		return
	}
	t, err = CreateTagIfNotExist(session, req.TagName, req.TagId)

	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
	}

	b.Title = req.Title
	b.Auther = req.Auther
	b.Description = req.Description
	b.TagId = t.TagId
	b.CreatedAt = time.Now()
	b.Icon = req.Icon
	if err != nil {
		ginrunner.ResponseJSON(ctx, errors.New("icon base64 cannot decoded"), nil)
		return
	}

	_, err = session.Table(b.TableName()).Insert(&b)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	content.BlogId = b.Id
	content.Content = req.Content

	if _, err := session.Table(content.TableName()).Insert(&content); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	session.Commit()
	var rsp BlogRsp
	rsp.Title = req.Title
	rsp.Auther = req.Auther
	rsp.Description = req.Description
	rsp.Tag = t.TagName
	rsp.TagId = t.TagId
	rsp.Icon = req.Icon
	rsp.CreatedAt = b.CreatedAt
	rsp.Id = b.Id
	server_context.SrvCtx.Logger.Infof("Blog created with id %d", b.Id)

	ginrunner.ResponseJSON(ctx, nil, rsp)
}

// Update Blog
// @Summary      Update Blog
// @Description  Update Blog
// @Tags         Blogs
// @Accept       json
// @Param        UpdateBlogReq   body      UpdateBlogReq  true  "UpdateBlogReq"
// @Produce      json
// @Success      200  {object}  BlogRsp
// @Router       /blogs [put]
func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	var req UpdateBlogReq
	var blog model.Blog
	var content model.Content
	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()

	if err := ctx.BindJSON(&req); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	if req.Id <= 0 {
		ginrunner.ResponseJSON(ctx, errors.New("invalid blog id"), nil)
	}
	has, err := session.Table(blog.TableName()).Where("id = ?", req.Id).Get(&blog)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	if !has {
		ginrunner.ResponseJSON(ctx, errors.New("blog id not found"), nil)
		return
	}

	if req.Title != "" {
		blog.Title = req.Title
	}
	if req.Auther != "" {
		blog.Auther = req.Auther
	}
	if req.Description != "" {
		blog.Description = req.Description
	}
	if req.Icon != "" {
		blog.Icon = req.Icon

	}

	if req.TagId > 0 || req.TagName != "" {
		t, err := CreateTagIfNotExist(session, req.TagName, req.TagId)
		blog.TagId = t.TagId
		if err != nil {
			ginrunner.ResponseJSON(ctx, err, nil)
			return
		}
	}

	_, err = session.Table(blog.TableName()).Where("id = ?", req.Id).Update(&blog)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	has, err = session.Table(content.TableName()).Where("blog_id = ?", req.Id).Get(&content)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	if !has {
		ginrunner.ResponseJSON(ctx, errors.New("blog id not found"), nil)
		return
	}
	if req.Content != "" {
		content.Content = req.Content
		_, err = session.Table(content.TableName()).Where("blog_id = ?", req.Id).Update(&content)
		if err != nil {
			ginrunner.ResponseJSON(ctx, err, nil)
			return
		}
	}

	session.Commit()
	ginrunner.ResponseJSON(ctx, nil, blog)
}

// Update Blog
// @Summary      Delete Blog
// @Description  Delete Blog
// @Tags         Blogs
// @Accept       json
// @Param        DeleteBlogReq   body      DeleteBlogReq  true  "DeleteBlogReq"
// @Produce      json
// @Success      200  {object}  DeleteBlogRsp
// @Router       /admin/blogs [delete]
func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	var req DeleteBlogReq
	var rsp DeleteBlogRsp
	var err error
	var b model.Blog

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	if err = ctx.BindJSON(&req); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	if req.Id <= 0 {
		ginrunner.ResponseJSON(ctx, errors.New("invalid blog id"), nil)
		return
	}

	_, err = session.Table(b.TableName()).Where("id = ?", req.Id).Delete(&b)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	if err := session.Commit(); err != nil {
		ginrunner.ResponseJSON(ctx, errors.New("fail to commit"), nil)
		return
	}
	rsp.Id = req.Id
	ginrunner.ResponseJSON(ctx, nil, rsp)
}
