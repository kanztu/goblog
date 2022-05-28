package model

type Content struct {
	ContentId int64  `xorm:"pk autoincr"`
	BlogId    int64  `xorm:"blog_id"`
	Content   string `xorm:"content size:67107840"`
}

func (nw *Content) TableName() string {
	return "content"
}
