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

func CreateTagIfNotExist(session *xorm.Session, tagName string, tagId int64) (int64, error) {
	var tag_id int64
	var t model.Tag
	if tagId > 0 {
		// Have Tag Id, use Tag first
		// check tagID exist
		has, err := session.Table(t.TableName()).Where("tag_id = ?", tagId).Get(&t)
		if err != nil {
			return 0, err
		}
		if !has {
			return 0, errors.New("tag not found")
		}
		tag_id = tagId
	} else if tagName != "" {
		// Have Tag name
		// check tag exist
		has, err := session.Table(t.TableName()).Where("name = ?", tagName).Get(&t)
		if err != nil {
			return 0, err
		}
		if has {
			tag_id = tagId
		} else {
			// Tag not found, create it
			var new_tag model.Tag
			new_tag.TagName = tagName
			id, err := session.Table(new_tag.TableName()).Insert(&new_tag)
			if err != nil {
				return 0, err
			}
			tag_id = id
		}
	}
	return tag_id, nil
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
