package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/pkg/ginrunner"
	"github.com/kanztu/goblog/pkg/model"
	"github.com/kanztu/goblog/pkg/server_context"
	"github.com/kanztu/goblog/pkg/utils"
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
	fname := htmlroot + "single.html"
	b, err := utils.ReadFileToByte(fname)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	ctx.Data(http.StatusOK, ginrunner.ContentTypeHTML, b)
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
