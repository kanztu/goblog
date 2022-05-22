package model

import "time"

type Blog struct {
	Id          int64  `xorm:"pk autoincr"`
	Title       string `xorm:"title"`
	Auther      string `xorm:"auther"`
	TagId       int64  `xorm:"index tag_id"`
	Icon        string `xorm:"icon"`
	Description string `xorm:"description"`

	CreatedAt time.Time `xorm:"created_at"`
}

type BlogTag struct {
	Blog `xorm:"extends"`
	Tag  `xorm:"extends"`
}

func (nw *Blog) TableName() string {
	return "blog"
}
