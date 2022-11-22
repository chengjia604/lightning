package jsfind

import (
	"blot/blot"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var html_data string

func Ordinary(B *blot.Ba) {
	//普通提取
	B.Scan(&html_data)
	Depth(B)

}

var wait sync.WaitGroup

func Depth(B *blot.Ba) {
	//深度提取
	fmt.Println("深度提取")

	//data := B.Html_url(html_data)
	//Start_data <- data
	//wait.Add(1)
	//wait.Wait()
}

//func manger(B *blot.Ba) {
//	fmt.Println("manger")
//	for {
//		select {
//		case url1 := <-url:
//			go func() {
//				//50个并发请求
//				B.Url = url1
//				B.Get().Scan(&de)
//				ch_data <- de
//			}()
//		case th_html := <-ch_data:
//			go func() {
//				data := B.Html_url(th_html)
//				js_context(B, data)
//			}()
//		case data := <-Start_data:
//			for _, url_data := range data {
//				if B.Domain(url_data) == B.DomainName { //判断域名
//					url <- url_data
//				}
//			}
//		}
//	}
//}

func js_context(B *blot.Ba, url_data []string) {

	//fuzz := config.Read_fuzz()
	for _, data := range url_data {
		if strings.HasPrefix(data, "http") || strings.HasPrefix(data, "https") {
			//判断开头
			continue
		} else {
			//提取内容
			var js_data string
			B.Url = B.Url + data
			B.Get().Scan(&js_data)
			data_separate(B,js_data)//js内容


			//for _, impression := range fuzz {
			//	if ok, _ := regexp.MatchString(".*"+impression+".*", js_data); ok {
			//		fmt.Println("包含铭感字符:", impression)
			//	}
			//}
		}
	}
}
func Js_path(context string) []string {
	recom, _ := regexp.Compile("(?<='|\")/.+?(?='|\")")
	path_data := recom.FindAllString(context, -1)
	return path_data
}

func data_separate(B *blot.Ba,context string) {
	//js和path分离
	for _, data := range Js_path(context) {
		if ok, _ := regexp.MatchString("(?<=)\\.(js)", data); ok {
			//js文件
			B.Url=
		} else if ok, _ := regexp.MatchString("(?<=)\\.(css)", data); ok {
			continue
		} else {
			//连接
		}

	}
}
func Route_extraction() {
	//提取js中的路由
}
