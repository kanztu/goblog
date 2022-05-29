package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/kanztu/goblog/pkg/ginrunner"
	"github.com/kanztu/goblog/pkg/model"
	"github.com/kanztu/goblog/pkg/server_context"
)

type TagController struct {
}

func NewTagController() *TagController {
	return &TagController{}
}

func CreateTagIfNotExist(session *xorm.Session, tagName string, tagId int64) (model.Tag, error) {
	var t model.Tag
	server_context.SrvCtx.Logger.Debug(tagId)
	server_context.SrvCtx.Logger.Debug(tagName)
	if tagId > 0 {
		// Have Tag Id, use Tag first
		// check tagID exist
		has, err := session.Table(t.TableName()).Where("tag_id = ?", tagId).Get(&t)
		if err != nil {
			return t, err
		}
		if !has {
			return t, errors.New("tag not found")
		}
	} else if tagName != "" {
		// Have Tag name
		// check tag exist
		has, err := session.Table(t.TableName()).Where("name = ?", tagName).Get(&t)
		server_context.SrvCtx.Logger.Debugf("has: %v", has)
		if err != nil {
			return t, err
		}
		if !has {
			// Tag not found, create it
			var new_tag model.Tag
			new_tag.TagName = tagName
			_, err := session.Table(new_tag.TableName()).Insert(&new_tag)
			if err != nil {
				return t, err
			}
			t = new_tag
		}
	}
	return t, nil
}

func (c *TagController) GetTag(ctx *gin.Context) {

	var rsp []GetTagRsp
	var tags []model.Tag

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	session.Table("tag").Find(&tags)
	for _, v := range tags {
		var r GetTagRsp
		r.TagId = v.TagId
		r.TagName = v.TagName
		rsp = append(rsp, r)
	}
	ginrunner.ResponseJSON(ctx, nil, rsp)
}

func (c *TagController) CreateTag(ctx *gin.Context) {
	var req CreateTagReq
	var rsp CreateTagRsp
	var tag model.Tag
	if err := ctx.BindJSON(&req); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	if req.TagName == "" {
		ginrunner.ResponseJSON(ctx, errors.New("tagname cannot be empty"), nil)
		return
	}
	tag.TagName = req.TagName

	session := server_context.SrvCtx.DB.NewSession()
	if _, err := session.Table(tag.TableName()).Insert(&tag); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	if err := session.Commit(); err != nil {
		ginrunner.ResponseJSON(ctx, errors.New("fail to commit"), nil)
		return
	}

	rsp.TagId = tag.TagId
	rsp.TagName = tag.TagName
	ginrunner.ResponseJSON(ctx, nil, rsp)
}

func (c *TagController) DeleteTag(ctx *gin.Context) {
	var req DeleteTagReq
	var rsp DeleteTagRsp
	var err error
	var t model.Tag

	session := server_context.SrvCtx.DB.NewSession()
	defer session.Close()
	if err = ctx.BindJSON(&req); err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}

	if req.Tagid <= 0 {
		ginrunner.ResponseJSON(ctx, errors.New("invalid blog id"), nil)
		return
	}

	_, err = session.Table(t.TableName()).Where("tag_id = ?", req.Tagid).Delete(&t)
	if err != nil {
		ginrunner.ResponseJSON(ctx, err, nil)
		return
	}
	if err := session.Commit(); err != nil {
		ginrunner.ResponseJSON(ctx, errors.New("fail to commit"), nil)
		return
	}
	rsp.Id = int64(req.Tagid)
	ginrunner.ResponseJSON(ctx, nil, rsp)
}
