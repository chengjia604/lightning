package main

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Ba struct {
	ERROR     error
	Resp_Data []byte
}

var resp *http.Response
var err error

func (B *Ba) Get(url string) (dt *Ba) {
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
	respstring, _ := io.ReadAll(resp.Body)
	B.Resp_Data = respstring
	dt = B
	return
}
func (B *Ba) Post(url string) {
	//post请求，返回[]byte
	if resp, err = http.Post(url, "application/json", strings.NewReader("username=test&password=ab123123")); err != nil {
		panic("post请求失败")
	}

}
func (B *Ba) scan(value any) {
	//将结果返回
	data_type := reflect.TypeOf(value)
	if data_type.Kind() == reflect.Ptr { //通过kind函数获取到是否为指针
		data_value := reflect.ValueOf(value).Elem()
		switch data_type.Elem().Kind() { //获取传入变量的类型种类
		case reflect.String:
			data_value.SetString(string(B.Resp_Data))
		case reflect.Struct:
			for i := 0; i < data_type.Elem().NumField(); i++ { //通过反射对结构体赋值
				name := B.json_assert()[data_type.Elem().Field(i).Name]
				if data_value.Type() == reflect.TypeOf(name) {
					data_value.FieldByName(data_type.Elem().Field(i).Name).Set(reflect.ValueOf(name))
				} else {
					panic("类型不匹配")
				}
			}
		}
	} else {
		panic("传入类型需要指针")
	}
}

func (B *Ba) json_assert() map[any]any {
	//判断是否为json
	var map_d map[any]any
	if err := json.Unmarshal(B.Resp_Data, &map_d); err != nil {
		panic("返回数据非json")
	}
	return map_d
}
func Start() (ba *Ba) {
	//项目入口
	ba = &Ba{}
	return
}
func main() {
	var a string
	Start().Get("http://www.baidu.com").scan(&a)
}
