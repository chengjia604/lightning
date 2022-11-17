package jsfind

import (
	"blot/blot"
	"blot/config"
	"fmt"
	"regexp"
	"strings"
)

func Ordinary(B *blot.Ba) {
	//普通提取
	B.Html_url(B.Get_data)

}

func depth() {
	//深度提取

}

func js_context(B *blot.Ba, url_data []string) {
	//js敏感内容提取
	for _, data := range url_data {
		if strings.HasPrefix(data, "http") || strings.HasPrefix(data, "https") {
			//判断开头
			continue
		} else {
			//提取内容
			var js_data string
			B.Url = B.Url + data
			B.Get().Scan(&js_data)
			for _, impression := range config.Read_fuzz() {
				if ok, _ := regexp.MatchString(".*"+impression+".*", js_data); ok {
					fmt.Println("包含铭感字符:", impression)
				}
			}
		}
	}
}
