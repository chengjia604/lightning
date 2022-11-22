package jsfind

import (
	"blot/blot"
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
	"sync"
)

var ordin = make(chan string, 30)
var l sync.Mutex
var w sync.WaitGroup

func Ordinary(B *blot.Ba) {
	//普通提取
	var html_data string
	B.Scan(&html_data)
	js_context(B, B.Html_url(html_data))
	fmt.Println(B.Html_url(html_data))
	go go_th(B)
	w.Add(1)
	w.Wait()
}

func Depth(B *blot.Ba) {
	//深度提取
	fmt.Println("深度提取")

}

func js_context(B *blot.Ba, url_data []string) {

	//fuzz := config.Read_fuzz()
	for _, data := range url_data {
		if strings.HasPrefix(data, "http") || strings.HasPrefix(data, "https") {
			//判断开头
			continue
		} else {
			//提取内容
			ordin <- data

			//for _, impression := range fuzz {
			//	if ok, _ := regexp.MatchString(".*"+impression+".*", js_data); ok {
			//		fmt.Println("包含铭感字符:", impression)
			//	}
			//}
		}
	}
}
func go_th(B *blot.Ba) {
	//多线程
	for {
		select {
		case js_data := <-ordin:
			go requ_js(B, js_data)
		}
	}
}
func requ_js(B *blot.Ba, data string) {
	l.Lock()
	defer l.Unlock()
	var js_data string
	B.Get(B.DomainName + data).Scan(&js_data)
	data_separate(js_data)

}

func Js_path(context string) []string {
	var path_data []string
	recom := regexp2.MustCompile("(?<='|\")/[a-zA-Z0-9]+?(?='|\")", 0)
	m, _ := recom.FindStringMatch(context)
	for m != nil {
		path_data = append(path_data, m.String())
		m, _ = recom.FindNextMatch(m)
	}
	return path_data
}

func data_separate(context string) {
	//js和path分离

	for _, data := range Js_path(context) {
		re := regexp2.MustCompile("(?<=)\\\\.(js)", 0)
		if ok, _ := re.MatchString(data); ok {
			fmt.Println(data)
			ordin <- data

		}
		//if ok, _ := regexp2.MustCompile("(?<=)\\.(js)", data); ok {
		//	//js文件
		//	fmt.Println(data)
		//	ordin <- data
		//} else if ok, _ := regexp.MatchString("(?<=)\\.(css)", data); ok {
		//	continue
		//} else {
		//	//连接
		//	fmt.Println(data)
		//}
	}
}
func Route_extraction() {
	//提取js中的路由
}
