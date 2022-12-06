package blot

import (
	"blot/structural"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

type Ba struct {
	ERROR      error
	RespData   []byte
	Json       map[string]any
	Ip         string
	Url        string
	DomainName string
	Subdom     string
	Regular
}

var L sync.Mutex
var Cookie string
var I string
var S string

func (B *Ba) Get(url string) *Ba {
	/*
		默认百度请求头，后期可通过命令设置
	*/
	if B.Url == "" {
		domain := strings.Split(url, "/")
		B.Url = domain[0] + "//" + domain[2]
		B.Subdom = strings.Split(domain[2], ".")[1]
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			panic("get关闭发生错误")
		}
	}(resp.Body)
	resp.Header.Set("user-agent", fmt.Sprintf("%s", structural.Useraget))
	resp.Header.Set("Accept", "*/*")
	//resp.Header.Add("cookie", Cookie)
	respstring, _ := io.ReadAll(resp.Body)
	B.RespData = respstring

	return B
}

func (B *Ba) PostJson(url string, json_data map[string]any) *Ba {
	/*
		默认百度请求头和json的格式
	*/
	//if B.DomainName == "" {
	//	B.DomainName = B.Domain(url)
	//}
	B.Json = json_data
	resp, err := http.Post(url, "", strings.NewReader(B.jsonData()))
	if err != nil {
		panic("post请求失败")
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	resp.Header.Set("content-type", "")
	resp.Header.Set("user-agent", "")

	respbyte, _ := io.ReadAll(resp.Body)
	B.RespData = respbyte
	//B.Post_data = string(respbyte)
	return B
}

func (B *Ba) Scan(value any) {
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
		panic(err)
	}
	return string(data)
}
func Start() (B *Ba) {
	//项目入口
	B = &Ba{}

	return
}
