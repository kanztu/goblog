package main

import "syscall/js"

func searchBlogByTagJsFunc(this js.Value, inputs []js.Value) interface{} {
	id := inputs[0].Int()
	searchBlogByTag(int64(id))
	return nil
}
