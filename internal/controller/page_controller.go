package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanztu/goblog/pkg/ginrunner"
	"github.com/kanztu/goblog/pkg/model"
	"github.com/kanztu/goblog/pkg/server_context"
	"github.com/kanztu/goblog/pkg/utils"
)

var (
	htmlroot = "src/html/"
)

type PageController struct {
}

func NewPageController() *PageController {
	return &PageController{}
}

func (c *PageController) GetIndexPage(ctx *gin.Context) {
	fname := htmlroot + "index.html"
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
