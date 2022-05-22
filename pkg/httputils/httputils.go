package httputils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func CallApi(Url string, method string, params interface{}) (resBody string, err error) {
	var payload io.Reader
	var req *http.Request
	transport := &http.Transport{}

	client := &http.Client{
		Transport: transport,
	}
	if params != nil {
		json_str, _ := params.([]byte)
		payload = strings.NewReader(string(json_str))
		req, err = http.NewRequest(method, Url, payload)
	} else {
		req, err = http.NewRequest(method, Url, nil)
	}

	// if method == http.MethodPost {
	// 	json_str, _ := params.([]byte)
	// 	payload = strings.NewReader(string(json_str))
	// 	req, err = http.NewRequest(method, Url, payload)
	// } else if method == http.MethodGet {
	// 	req, err = http.NewRequest(method, Url, nil)
	// }

	if err != nil {
		err = fmt.Errorf("callApi: client.Do err: %s", err)
		return
	}

	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Accept", "application/json")

	if method == http.MethodGet || method == http.MethodPost {
		p, _ := params.(url.Values)
		if p != nil {
			req.URL.RawQuery = p.Encode()
		}
	}

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("CallApi: client.Do err: %s", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("CallApi: ioutil.ReadAll err: %s", err)
		return
	}

	//fmt.Printf("CallApi body: %s\n\n", string(body))

	return string(body), nil
}
