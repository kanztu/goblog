package main

import (
	"fmt"

	"github.com/kanztu/goblog/internal/controller"
	"github.com/kanztu/goblog/pkg/wasm"
)

func fetchPageCata() {
	var rsp []controller.GetPageCataRsp
	wasm.WasmReq(wasm.GetHostUrl()+"/api/v1/pagecata", &rsp, nil)
	l.Info(len(rsp))

	bodyDOM := wasm.GetElementById("page_cata")
	bodyDOM.DeleteInnerHTML()

	for _, v := range rsp {
		isActive := ""
		if wasm.GetPath() == "/blog"+v.CataPath {
			isActive = "colorlib-active"
		}
		bodyDOM.AppendInnerHTML(fmt.Sprintf(cataHtml, isActive, v.CataPath, v.CataName))
	}
}

func fetchTagCloud() {
	var rsp []controller.GetTagRsp
	wasm.WasmReq(wasm.GetHostUrl()+"/api/v1/tags", &rsp, nil)
	l.Info(len(rsp))

	bodyDOM := wasm.GetElementById("tagcloud")
	bodyDOM.DeleteInnerHTML()
	for _, v := range rsp {
		bodyDOM.AppendInnerHTML(fmt.Sprintf(tag_cloud_html, v.TagId, v.TagName))
	}
}

func searchBlogByTag(id int64) {
	var rsp []controller.BlogRsp
	var req controller.GetBlogReq
	req.TagId = id

	wasm.WasmReq(wasm.GetHostUrl()+"/api/v1/blogs/tag", &rsp, req)
	l.Info(len(rsp))
	bodyDOM := wasm.GetElementById("blog_list")
	bodyDOM.DeleteInnerHTML()
	for _, v := range rsp {
		url := fmt.Sprintf("/id/%d", v.Id)
		bodyDOM.AppendInnerHTML(fmt.Sprintf(blog_preview_html, url, v.Icon, url, v.Title, v.CreatedAt, v.TagId, v.Tag, v.Description, url))
	}
}
