package main

import (
	"fmt"

	"github.com/kanztu/goblog/internal/controller"
	"github.com/kanztu/goblog/pkg/wasm"
)

func fetchBlog() {
	var rsp []controller.BlogRsp
	wasm.WasmReq(wasm.GetHostUrl()+"/api/v1/blogs", &rsp, nil)
	l.Info(len(rsp))
	bodyDOM := wasm.GetElementById("blog_list")
	bodyDOM.DeleteInnerHTML()
	for _, v := range rsp {
		url := fmt.Sprintf("/id/%d", v.Id)
		bodyDOM.AppendInnerHTML(fmt.Sprintf(blog_preview_html, url, v.Icon, url, v.Title, v.CreatedAt, v.TagId, v.Tag, v.Description, url))
	}
}
