package main

import (
	"github.com/kanztu/goblog/pkg/logger"
	"github.com/kanztu/goblog/pkg/wasm"
	"github.com/sirupsen/logrus"
)

var l *logrus.Entry

func main() {
	c := make(chan bool)

	l = logger.InitLogger(logger.LEVEL_INFO, "wasm")
	if wasm.GetPath() == "/blog/about" {
		Animation()
	}

	<-c
}
