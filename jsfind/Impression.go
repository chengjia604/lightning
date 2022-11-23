package jsfind

import (
	"blot/blot"
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
	"sync"
)

var (
	ordin = make(chan string, 30)
	dep   = make(chan string, 30)
)
var l sync.Mutex
var w sync.WaitGroup
var url []string
var js []string

func Ordinary(B *blot.Ba) {
	//普通提取
	var html_data string

	B.Scan(&html_data)

	go go_th(B)
	js_context(B.Html_url(html_data), "Ordinary")
	w.Add(1)
	w.Wait()
}

func Depth(B *blot.Ba) {
	//深度提取
	fmt.Println("深度提取")
}

func js_context(url_data []string, typename string) {

	//fuzz := config.Read_fuzz()

	for _, data := range url_data {
		if strings.HasPrefix(data, "http") || strings.HasPrefix(data, "https") {
			//http连接
			if typename == "Ordinary" {
				continue
			} else {

			}
		} else {
			//.js连接
			if typename == "Ordinary" {
				fmt.Println(data)
				ordin <- data
			} else {
				dep <- data
			}

			//for _, impression := range fuzz {
			//	if ok, _ := regexp.MatchString(".*"+impression+".*", js_data); ok {
			//		fmt.Println("包含铭感字符:", impression)
			//	}
			//}
		}
	}

}

func go_th(B *blot.Ba) {
	//监听数据，
	fmt.Println(111)
	for {
		select {
		case js_data := <-ordin:
			go requ_js(B, js_data, "ord")
		case js_data1 := <-dep:
			go requ_js(B, js_data1, "de")
		
		}
	}
}

func requ_js(B *blot.Ba, data string, typename string) {
	l.Lock()
	defer l.Unlock()
	var js_data string
	B.Get(B.DomainName + data).Scan(&js_data)
	data_separate(js_data, typename)
}

func Js_path(context string) []string {
	var path_data []string
	//提取js文件中的.js和/xxx路径
	recom := regexp2.MustCompile("(?<='|\")/[a-zA-Z].+?(?='|\")", 0) //(?<='|")/[a-zA-Z].+?(?='|")
	m, _ := recom.FindStringMatch(context)
	for m != nil {
		path_data = append(path_data, m.String())
		m, _ = recom.FindNextMatch(m)
	}
	return path_data
}

var (
	rejs  = regexp2.MustCompile("(?<=)\\.(js)", 0)
	recss = regexp2.MustCompile("(?<=)\\.(css)", 0)
)

func data_separate(context string, typename string) {
	//js和path分离

	for _, data := range Js_path(context) {
		if ok, _ := rejs.MatchString(data); ok {
			//js文件继续请求内容
			if typename == "ord" {
				js = append(js, data)
			} else {
				dep <- data
			}
		} else if ok, _ := recss.MatchString(data); ok {
			continue
		} else {
			// /xxx的链接
			if typename == "ord" {
				url = append(url, data)
			}
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
