package jsfind

import (
	"blot/blot"
	"blot/config"
	"blot/structural"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/fatih/color"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var (
	ord               = make(chan string, 30)
	w                 sync.WaitGroup
	B                 *blot.Ba
	l                 sync.Mutex
	jsurl             = make(map[string]bool)
	url               = make(map[string]bool)
	httpurl           = make(map[string]bool)
	funzz             = config.Read_fuzz()
	sensitive_url     = make(map[string][]string)
	sensitive_domname = make(map[string][]string)
	//url []string
)

func Ord(b *blot.Ba) {
	var html string

	B = b
	B.Scan(&html)
	go go_th()
	html_map := B.Html_url(html)

	w.Add(1)
	http_js(html_map)
	w.Wait()

	if blot.I != "" {
		config.Create_html(sensitive_url, sensitive_domname, jsurl, blot.I, B.Url)
	}
	if blot.S {
		w.Add(1)
		go fturl()
		w.Wait()
	}

	color.Green("收集完成！！！")
}
func go_th() {
	for {
		select {
		case data := <-ord:
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
			if strings.Split(k, "/")[2] == B.Subdom {
				url[k] = true
			} else {
				httpurl[k] = true
			}
		} else {
			ord <- k
			w.Add(1)
		}
	}
}

func js_requ(data string) {
	defer w.Done()
	//var js_context string
	//B.Get(B.Url + data).Scan(&js_context)
	l.Lock()
	jsurl[data] = true
	l.Unlock()
	var js_context []byte
	resp, err := http.Get(B.Url + data)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			panic("get关闭发生错误")
		}
	}(resp.Body)

	resp.Header.Set("user-agent", structural.Useraget)
	resp.Header.Set("Accept", "*/*")
	resp.Header.Add("cookie", blot.Cookie)
	js_context, _ = io.ReadAll(resp.Body)
	url_js(string(js_context))
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
	rejs   = regexp2.MustCompile("/*.js+?(?=\"|'|>)", 0)
	recss  = regexp2.MustCompile("(?<=)\\.(css)", 0)
	rehttp = regexp2.MustCompile("(?<=https?://)[^/].+?(?=/|\"|')", 0) //提取域名
)

func url_js(conext string) {

	for _, value := range jscontext(conext) {

		if ok, _ := rejs.MatchString(value); ok {
			ord <- value
		} else if strings.HasPrefix(value, "https") || strings.HasPrefix(value, "http") {
			if strings.Split(value, "/")[2] == B.Subdom {
				l.Lock()
				url[value] = true
				l.Unlock()
			} else {
				w.Add(1)
				go fuzz(value, 2)
			}
		} else {
			//blot.L.Lock()
			//url[value] = true
			//blot.L.Unlock()
			w.Add(1)
			go fuzz(value, 1)
		}

	}
}
func fturl() {
	//url
	defer w.Done()
	fmt.Println("目标资产:")
	for k, _ := range sensitive_url {
		color.Green(k + "\n")
	}
	w.Add(1)
	go domname()
}

func domname() {
	//domname
	defer w.Done()
	fmt.Println("其他域名资产:")
	for k, _ := range sensitive_domname {
		color.Green(k + "\n")
	}
	w.Add(1)
	go ftjs()
}
func ftjs() {
	//js
	defer w.Done()
	fmt.Println("js资产:")
	for k, _ := range jsurl {
		color.Green(k + "\n")
	}
}

func fuzz(url string, pandaun int) {
	//匹配敏感字符
	defer w.Done()
	var sensitive []string
	for _, value := range funzz {
		if ok, _ := regexp.MatchString(".*"+value+".*", url); ok {
			sensitive = append(sensitive, value)
		}
	}

	if pandaun == 1 {
		l.Lock()
		sensitive_url[url] = sensitive
		l.Unlock()
	} else if pandaun == 2 {
		l.Lock()
		sensitive_domname[url] = sensitive
		l.Unlock()
	}

}
