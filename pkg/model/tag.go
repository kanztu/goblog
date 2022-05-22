package model

type Tag struct {
	TagId   int64  `xorm:"pk autoincr"`
	TagName string `xorm:"name"`
}

func (nw *Tag) TableName() string {
	return "tag"
}
