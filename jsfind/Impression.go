package jsfind

import (
	"blot/blot"
	"blot/config"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/fatih/color"
	"regexp"
	"strings"
	"sync"
)

var (
	ordin = make(chan string, 30)
	dep   = make(chan string, 50)
)

var w sync.WaitGroup
var url []string
var js1 = make(map[string][]string)
var B *blot.Ba

func Ordinary(b *blot.Ba) {
	//普通提取
	var html_data string
	B = b
	B.Scan(&html_data)
	go go_th()
	go js_context(B.Html_url(html_data), "Ordinary")
	w.Add(1)
	w.Wait()
	w.Add(1)
	go js_parse()
	w.Wait()

}

func js_parse() {
	//解析js
	defer w.Done()
	for k, value := range js1 {
		color.Green(k, value)
	}
	w.Wait()
}
func Depth(B *blot.Ba) {
	//深度提取
	fmt.Println("深度提取")
}

var SubdomainName = make(map[string]bool) //子域名
func js_context(url_data map[string]bool, typename string) {
	//爬取第二层应该是用一个线程，而不是一个http一个线程
	defer w.Done()
	for k, _ := range url_data {
		if strings.HasPrefix(k, "http") || strings.HasPrefix(k, "https") {
			//http连接
			//if typename == "Ordinary" {
			//	url = append(url, k)
			//	continue
			//} else {
			//	if B.Domain(k) == B.DomainName {
			//		http_data = append(http_data, k)
			//	}
			//}
			a, b := B.Subdomain(k)
			if B.Subdom == a {
				url = append(url, k)
			} else {
				if b == B.DomainName {
					SubdomainName[k] = true
					//dep <- k //子域名
				}
			}
		} else {
			//.js连接
			w.Add(1)
			ordin <- k
		}
	}
}

func requ_http(http []string, i int) {
	//批量处理http

}

func go_th() {
	//监听数据，

	for {
		select {
		case js_data := <-ordin: //.js
			go requ_js(js_data)
		case js_data1 := <-dep:
			go requ_js(js_data1)
		}
	}
}

func requ_js(data string) {
	defer w.Done()
	blot.L.Lock()
	var js_data string
	defer blot.L.Unlock()
	B.Get(B.Url + data).Scan(&js_data)
	//go fuzz(data, js_data)
	data_separate(js_data)
}
func Js_path(context string) []string {
	var path_data []string
	//提取js文件中的.js和/xxx路径
	recom := regexp2.MustCompile("(?<='|\")/[a-zA-Z].+?(?='|\")|(https?|http|ftp|file):\\/\\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]", 0) //(?<='|")/[a-zA-Z].+?(?='|")
	m, _ := recom.FindStringMatch(context)
	for m != nil {
		path_data = append(path_data, m.String())
		m, _ = recom.FindNextMatch(m)
	}

	return path_data
}

var (
	rejs   = regexp2.MustCompile("/*.js+?(?=\"|')", 0)
	recss  = regexp2.MustCompile("(?<=)\\.(css)", 0)
	rehttp = regexp2.MustCompile("(?<=https?://)[^/].+?(?=/|\"|')", 0) //提取域名
)

func data_separate(context string) {
	//js和path分离
	for _, data := range Js_path(context) {
		if ok, _ := rejs.MatchString(data); ok {
			//js文件继续请求内容
			//if typename == "ord" {
			//	js1 = append(js1, data)
			//} else {
			//	dep <- data
			//}
			fmt.Println(data)
			ordin <- data
		} else if okcss, _ := recss.MatchString(data); okcss {
			continue
		} else if okhttp, _ := rehttp.MatchString(data); okhttp {
			a, b := B.Subdomain(data)
			if B.Subdom == a {
				url = append(url, data)
			} else {
				if B.DomainName == b {
					SubdomainName[data] = true
				}
			}
		} else {
			// /xxx的链接
			if strings.HasPrefix(data, "http") || strings.HasPrefix(data, "https") {
				continue
			}
			url = append(url, data)
		}
	}
}

var fuzzcontext = config.Read_fuzz()
var minggan []string

func fuzz(data string, context string) {
	w.Add(1)
	defer w.Done()
	for _, impression := range fuzzcontext {
		if ok, _ := regexp.MatchString(".*"+impression+".*", context); ok {
			minggan = append(minggan, impression)
		}
	}
	js1[data] = minggan
}
