package main

import (
	"io"
	"net/http"
	"strings"
)

type Ba struct {
	ERROR error
	Data  string
}

var resp *http.Response
var err error

func (B *Ba) Get(url string) (dt *Ba) {
	//get请求

	if resp, err = http.Get(url); err != nil {
		panic("请求get失败")
	}
	defer resp.Body.Close()
	respstring, _ := io.ReadAll(resp.Body)
	B.Data = string(respstring)
	dt = B
	return
}
func (B *Ba) Post(url string) {
	//post请求
	if resp, err = http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("username=test&password=ab123123")); err != nil {
		panic("post请求失败")
	}
}

func (B *Ba) scan(value *string) {
	//将结果返回
	*value = B.Data
}
func Start() (ba *Ba) {
	//项目入口
	ba = &Ba{}
	return
}
