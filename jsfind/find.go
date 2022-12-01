package jsfind

import (
	"blot/blot"
	"github.com/dlclark/regexp2"
	"github.com/gookit/color"
	"strings"
	"sync"
)

var (
	ord   = make(chan string, 50)
	w     sync.WaitGroup
	B     *blot.Ba
	l     sync.Mutex
	jsurl = make(map[string]bool)
	//url   = make(map[string]bool)
	url []string
)

func Ord(b *blot.Ba) {
	var html string
	color.Greenf("数据收集中。。。。")
	B = b
	B.Scan(&html)
	go go_th()
	html_map := B.Html_url(html)
	w.Add(1)
	http_js(html_map)
	w.Wait()

	w.Add(1)
	go fturl()
	w.Wait()
}
func go_th() {
	for {
		select {
		case data := <-ord:
			w.Add(1)
			go js_requ(data)
		}
	}
}

func http_js(data map[string]bool) {
	defer w.Done()
	for k, _ := range data {
		http := strings.Split(k, ":")[0]
		https := strings.Split(k, ":")[0]
		if https == "https" || http == "http" {
			continue
		} else {
			ord <- k
		}
	}
}

var js_context string

func js_requ(data string) {
	defer w.Done()
	B.Get(B.Url + data).Scan(&js_context)
	url_js(js_context)
}

func jscontext(context string) []string {
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

func url_js(conext string) {

	for _, value := range jscontext(conext) {
		if ok, _ := rejs.MatchString(value); ok {
			ord <- value
		} else {
			//url[value] = true
			l.Lock()
			url = append(url, value)
			l.Unlock()
		}
	}
}

func fturl() {
	w.Done()
	for _, value := range url {
		color.Greenf(value + "\n")
	}
}
