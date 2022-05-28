package main

import "syscall/js"

func searchBlogByTagJsFunc(this js.Value, args []js.Value) interface{} {
	id := args[0].Int()
	searchBlogByTag(int64(id))
	return nil
}
