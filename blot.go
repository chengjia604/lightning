package main

import (
	"blot/base"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Ba struct {
	ERROR    error
	RespData []byte
	Json     string
}

var resp *http.Response
var err error

func (B *Ba) Get(url string) *Ba {
	//get请求
	if resp, err = http.Get(url); err != nil {
		panic("请求get失败")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic("get关闭发生错误")
		}
	}(resp.Body)

	resp.Header.Set("content-type", "text/html;charset=utf-8")

	respstring, _ := io.ReadAll(resp.Body)
	B.RespData = respstring
	return B
}

func (B *Ba) PostJson(url string) *Ba {
	//json请求，返回[]byte
	B.Json = url
	if resp, err = http.Post(url, "application/json", strings.NewReader(B.jsonData())); err != nil {
		panic("post请求失败")
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	respbyte, _ := io.ReadAll(resp.Body)
	B.RespData = respbyte
	return B
}

func (B *Ba) scan(value any) {
	//结构化赋值
	dataType := reflect.TypeOf(value)
	if dataType.Kind() == reflect.Ptr { //通过kind函数获取到是否为指针
		dataValue := reflect.ValueOf(value).Elem()
		switch dataType.Elem().Kind() { //获取传入变量的类型种类
		case reflect.String:
			dataValue.SetString(string(B.RespData))
		case reflect.Struct:
			for i := 0; i < dataType.Elem().NumField(); i++ { //通过反射对结构体赋值
				name := B.jsonAssert()[strings.ToLower(dataType.Elem().Field(i).Name)] //结构体名称转小写
				if dataValue.Type() == reflect.TypeOf(name) {
					dataValue.FieldByName(dataType.Elem().Field(i).Name).Set(reflect.ValueOf(name))
				} else {
					panic("类型不匹配")
				}
			}
		}
	} else {
		panic("传入类型需要指针")
	}
}

func (B *Ba) jsonAssert() map[any]any {
	//判断是否为json
	var mapD map[any]any
	if err := json.Unmarshal(B.RespData, &mapD); err != nil {
		panic("返回数据非json")
	}
	return mapD
}

func (B *Ba) jsonData() string {
	//转换成json数据
	data, err := json.Marshal(B.Json)
	if err != nil {
		panic("转化结构体失败")
	}
	return string(data)
}
func Start() (B *Ba) {
	//项目入口
	B = &Ba{}
	return
}
func main() {
	var a string
	Start().Get("http://www.diji-it.com/").scan(&a)
	base.Html_url(a)

}
