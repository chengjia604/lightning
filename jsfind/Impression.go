package jsfind

import (
	"blot/blot"
	"blot/config"
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

var ch_data = make(chan string, 50)
var url = make(chan string, 50)
var Start_data = make(chan []string)
var de string
var wait sync.WaitGroup

func Depth(B *blot.Ba) {
	//深度提取
	fmt.Println("深度提取")
	go manger(B)
	data := B.Html_url(html_data)
	Start_data <- data

	wait.Add(1)
	wait.Wait()
}

func manger(B *blot.Ba) {
	fmt.Println("manger")
	var a []string
	for {
		select {
		case url1 := <-url:
			go func() {
				//50个并发请求
				B.Url = url1
				B.Get().Scan(&de)
				ch_data <- de
			}()

		case th_html := <-ch_data:
			go func() {
				data := B.Html_url(th_html)
				js_context(B, data)
			}()
		case data := <-Start_data:
			for _, url_data := range data {
				if B.Domain(url_data) == B.DomainName { //判断域名
					a = append(a, url_data)
					b := false
					for _, j := range a {
						if url_data == j {
							b = true
						}
					}
					if b {
						continue
					}
					url <- url_data
				}
			}
		}
	}
}

func js_context(B *blot.Ba, url_data []string) {
	//js敏感内容提取
	fuzz := config.Read_fuzz()
	for _, data := range url_data {
		if strings.HasPrefix(data, "http") || strings.HasPrefix(data, "https") {
			//判断开头
			continue
		} else {
			//提取内容
			var js_data string
			B.Url = B.Url + data
			B.Get().Scan(&js_data)
			for _, impression := range fuzz {
				if ok, _ := regexp.MatchString(".*"+impression+".*", js_data); ok {
					fmt.Println("包含铭感字符:", impression)
				}
			}
		}
	}
}
func Route_extraction() {
	//提取js中的路由
}
