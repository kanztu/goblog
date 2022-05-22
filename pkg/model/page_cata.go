package model

type PageCata struct {
	CataId   int64  `xorm:"pk autoincr"`
	CataName string `xorm:"cata_name"`
	CataPath string `xorm:"cata_path"`
}

func (nw *PageCata) TableName() string {
	return "page_cata"
}
