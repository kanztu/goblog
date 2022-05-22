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

func (c *BlogController) GetBlog(ctx *gin.Context) {
	var rsp []BlogRsp
	var blogtags []model.BlogTag

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()

	sql := session.Table("blog").Join("INNER", "tag", "blog.tag_id = tag.tag_id")
	// if req.TagId >= 0 {
	// 	sql = sql.Where("blog.tag_id = ?", req.TagId)
	// }
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

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var err error
	var req CreateBlogReq
	var t model.Tag
	var b model.Blog
	var content model.Content
	var tag_id int64

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	if err = ctx.BindJSON(&req); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	if req.TagName == "" {
		ginrunner.ResponseJSON(ctx, errors.New("no tag provided"), nil)
		return
	}
	tag_id, err = CreateTagIfNotExist(session, req.TagName, req.TagId)

	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
	}

	b.Title = req.Title
	b.Auther = req.Auther
	b.Description = req.Description
	b.TagId = tag_id
	b.CreatedAt = time.Now()
	b.Icon = req.Icon
	if err != nil {
		ginrunner.ResponseJSON(ctx, errors.New("icon base64 cannot decoded"), nil)
		return
	}

	blog_id, err := session.Table(b.TableName()).Insert(&b)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	content.BlogId = blog_id
	content.Content = req.Content

	_, err = session.Table(content.TableName()).Insert(&content)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	session.Commit()
	var rsp BlogRsp
	rsp.Title = req.Title
	rsp.Auther = req.Auther
	rsp.Description = req.Description
	rsp.Tag = t.TagName
	rsp.CreatedAt = b.CreatedAt
	rsp.Id = blog_id

	ginrunner.ResponseJSON(ctx, nil, rsp)
}

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
		blog.TagId, err = CreateTagIfNotExist(session, req.TagName, req.TagId)
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

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	var req DelBlogReq
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
	session.Commit()
	ginrunner.ResponseJSON(ctx, nil, "OK")
}
