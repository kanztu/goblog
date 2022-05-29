package controller

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/pkg/config"
	"github.com/kanztu/goblog/pkg/ginrunner"
	"github.com/kanztu/goblog/pkg/md"
	"github.com/kanztu/goblog/pkg/model"
	"github.com/kanztu/goblog/pkg/server_context"
)

const (
	htmlroot = "template/"
)

type PageController struct {
}

func NewPageController() *PageController {
	return &PageController{}
}

func (c *PageController) GetIndexPage(ctx *gin.Context) {
	fname := "index.html"

	var pageCata []model.PageCata
	var blogtags []model.BlogTag
	var rsp []BlogRsp
	var p model.PageCata
	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	session.Table(p.TableName()).Find(&pageCata)

	session.Table("blog").Join("INNER", "tag", "blog.tag_id = tag.tag_id").Find(&blogtags)
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

	server_context.SrvCtx.Logger.Info(rsp)
	ctx.HTML(http.StatusOK, fname, gin.H{
		"title":       SiteTitle,
		"site_author": SiteAuthor,
		"cata":        pageCata,
		"blog":        rsp,
	})
}

func (c *PageController) GetBlogPage(ctx *gin.Context) {
	var pageCata []model.PageCata
	var p model.PageCata
	var b model.BlogTag
	var con model.Content
	id := ctx.Param("id")
	server_context.SrvCtx.Logger.Debug(id)
	fname := "blog.html"
	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	session.Table(p.TableName()).Find(&pageCata)
	session.Table("blog").Join("INNER", "tag", "blog.tag_id = tag.tag_id").Where("blog.id = ?", id).Get(&b)
	session.Table(con.TableName()).Where("blog_id = ?", b.Id).Get(&con)

	mdfname := filepath.Join(config.CfgGlobal.Stor.Blog, con.Content)
	server_context.SrvCtx.Logger.Debugf("Load markdown: %v", mdfname)
	content, err := md.FetchMDToHtml(mdfname)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	ctx.HTML(http.StatusOK, fname, gin.H{
		"title":       SiteTitle,
		"site_author": SiteAuthor,
		"cata":        pageCata,
		"BlogTitle":   b.Title,
		"Description": b.Description,
		"Content":     template.HTML(content),
	})
}

func (c *PageController) GetPageCata(ctx *gin.Context) {
	var rsp []GetPageCataRsp
	var pageCata []model.PageCata
	var p model.PageCata
	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	session.Table(p.TableName()).Find(&pageCata)

	for _, v := range pageCata {
		var p_tmp GetPageCataRsp
		p_tmp.Id = v.CataId
		p_tmp.CataName = v.CataName
		p_tmp.CataPath = v.CataPath
		rsp = append(rsp, p_tmp)
	}
	ginrunner.ResponseJSON(ctx, nil, rsp)
}
