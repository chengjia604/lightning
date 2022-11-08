package request

import (
	"io"
	"net/http"
)

func (B *Ba) Get(url string) {
	//get请求
	var resp *http.Response
	var err error
	if resp, err = http.Get(url); err != nil {
		panic("请求get失败")
	}
	defer resp.Body.Close()
	respstring, err := io.ReadAll(resp.Body)

}

func Post() {
	//post请求

}
