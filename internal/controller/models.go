package controller

import "time"

type BlogRsp struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Auther      string `json:"auther"`
	Tag         string `json:"tag"`
	TagId       int64  `json:"tag_id"`
	Icon        string `json:"icon"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
}

type GetBlogReq struct {
	Title     string    `json:"title"`
	TagId     int64     `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateBlogReq struct {
	Title       string `json:"title"`
	Auther      string `json:"auther"`
	TagId       int64  `json:"tag_id"`
	TagName     string `json:"tag_name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`

	Content string `json:"content"`
}

type UpdateBlogReq struct {
	Id int64 `json:"id"`
	CreateBlogReq
}

type DelBlogReq struct {
	Id int64 `json:"id"`
}

type GetPageCataRsp struct {
	Id       int64  `json:"id"`
	CataName string `json:"cata_name"`
	CataPath string `json:"cata_path"`
}

type GetTagRsp struct {
	TagId   int64  `json:"tag_id"`
	TagName string `json:"name"`
}
