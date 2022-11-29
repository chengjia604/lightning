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
	dep   = make(chan string, 30)
)

var w sync.WaitGroup
var url = make(map[string]bool)
var js1 = make(map[string]bool)
var B *blot.Ba

func Ordinary(b *blot.Ba) {
	//普通提取
	var html_data string
	B = b
	B.Scan(&html_data)
	go go_th()
	go js_context(B.Html_url(html_data), "ord")
	w.Add(1)
	w.Wait()
	w.Add(1)
	go url_parse()
	w.Wait()
}
func Depth(b *blot.Ba) {
	//深度提取
	var html_data string
	B = b
	B.Scan(&html_data)
	go go_th()
	go js_context(B.Html_url(html_data), "dep")
	w.Add(1)
	w.Wait()

}

func js_parse() {
	//解析js
	defer w.Done()
	fmt.Println("js资产:\n")
	for k, _ := range js1 {
		color.Green(fmt.Sprintf("%s", k))
	}

}
func url_parse() {
	//输出url
	w.Add(1)
	defer w.Done()
	fmt.Println("url资产:\n")
	for k, _ := range url {
		color.Green(fmt.Sprintf("%s%s", B.Url, k))
	}
	go subdomname()
}

func subdomname() {
	//子域名
	w.Add(1)
	defer w.Done()
	fmt.Println("子域名资产:\n")
	for K, _ := range SubdomainName {
		color.Green(fmt.Sprintf("%s", K))
	}
	go js_parse()
}

var SubdomainName = make(map[string]bool) //子域名
func js_context(url_data map[string]bool, typename string) {
	defer w.Done()
	for k, _ := range url_data {
		if strings.HasPrefix(k, "http") || strings.HasPrefix(k, "https") {
			//http连接
			w.Add(1)
			go subdom(k, typename)
		} else {
			//.js连接
			w.Add(1)
			ordin <- k
		}
	}

}
func subdom(k string, name string) {
	defer w.Done()
	blot.L.Lock()
	defer blot.L.Unlock()
	a, b := B.Subdomain(k)
	if B.Subdom == a {
		url[k] = true
	} else {
		if b == B.DomainName {
			SubdomainName[k] = true
			if name == "dep" {
				
				dep <- k //子域名链接
			}
		}
	}
}

func requ_http(http string) {
	//批量处理http
	blot.L.Lock()
	defer blot.L.Unlock()
	var html string
	B.Get(http).Scan(&html)                //子域名请求
	go js_context(B.Html_url(html), "dep") //请求第二层

}

func go_th() {
	//监听数据，
	for {
		select {
		case js_data := <-ordin: //.js
			go requ_js(js_data)
		case js_data1 := <-dep:
			go requ_http(js_data1)
		}
	}
}

func requ_js(data string) {
	defer w.Done()
	blot.L.Lock()
	var js_data string
	defer blot.L.Unlock()
	js1[data] = true
	B.Get(B.Url + data).Scan(&js_data)
	//go fuzz(data, js_data)
	go data_separate(js_data)
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
	w.Add(1)

	blot.L.Lock()
	defer w.Done()
	defer blot.L.Unlock()
	for _, data := range Js_path(context) {
		if ok, _ := rejs.MatchString(data); ok {
			ordin <- data
		} else if okcss, _ := recss.MatchString(data); okcss {
			continue
		} else if okhttp, _ := rehttp.MatchString(data); okhttp {
			a, b := B.Subdomain(data)
			if B.Subdom == a {
				url[data] = true
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
			url[data] = true
		}
	}
}

var fuzzcontext = config.Read_fuzz()

func fuzz(data string, context string) {
	w.Add(1)
	defer w.Done()

	var minggan []string
	for _, impression := range fuzzcontext {
		if ok, _ := regexp.MatchString(".*"+impression+".*", context); ok {
			minggan = append(minggan, impression)
		}
	}

}
