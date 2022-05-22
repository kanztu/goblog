package wasm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syscall/js"

	"github.com/kanztu/goblog/pkg/httputils"
)

type DOM struct {
	Dom js.Value
}

func WasmReq(apiurl string, rsp interface{}, params interface{}) {
	var body string
	var err error
	if params != nil {
		body, err = httputils.CallApi(apiurl, http.MethodGet, params)
	} else {
		body, err = httputils.CallApi(apiurl, http.MethodGet, nil)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal([]byte(body), &rsp)
	return
}

func GetHostUrl() string {
	locationDOM := js.Global().Get("location")
	hostname := locationDOM.Get("host")
	scheme := locationDOM.Get("protocol")
	return fmt.Sprintf("%s//%s", scheme, hostname)
}

func GetPath() string {
	locationDOM := js.Global().Get("location")
	return locationDOM.Get("pathname").String()
}

func GetElementById(id string) DOM {
	var d DOM
	doc := js.Global().Get("document")
	d.Dom = doc.Call("getElementById", id)
	return d
}

func (d DOM) GetInnerHTML() string {
	return d.Dom.Get("innerHTML").String()
}

func (d DOM) SetInnerHTML(html string) {
	d.SetInnerHTML(html)

}

func (d DOM) AppendInnerHTML(html string) {
	d.Dom.Set("innerHTML", d.GetInnerHTML()+html)
}

func (d DOM) DeleteInnerHTML() {
	d.Dom.Set("innerHTML", "")
}
