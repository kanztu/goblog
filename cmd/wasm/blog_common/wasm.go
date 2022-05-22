package main

import (
	"syscall/js"

	"github.com/kanztu/goblog/pkg/logger"
	"github.com/kanztu/goblog/pkg/wasm"
	"github.com/sirupsen/logrus"
)

var l *logrus.Entry

func main() {
	c := make(chan bool)

	l = logger.InitLogger(logger.LEVEL_INFO, "wasm")
	fetchPageCata()
	if wasm.GetPath() == "/blog/" {
		fetchBlog()
		fetchTagCloud()
	}

	js.Global().Set("searchBlogByTag", js.FuncOf(searchBlogByTagJsFunc))

	<-c
}
