package model

type Content struct {
	ContentId int64  `xorm:"pk autoincr"`
	BlogId    int64  `xorm:"blog_id"`
	Content   string `xorm:"content"`
}

func (nw *Content) TableName() string {
	return "content"
}
