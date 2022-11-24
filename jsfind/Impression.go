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
	dep   = make(chan string, 50)
)
var (
	url_recevie = make(chan []string)
)
var l sync.Mutex
var w sync.WaitGroup
var url []string
var js1 []string
var number_de = make(chan int, 3)

func Ordinary(B *blot.Ba) {
	//普通提取
	var html_data string

	B.Scan(&html_data)

	go go_th(B)
	go js_context(B, B.Html_url(html_data), "Ordinary")
	w.Add(1)
	w.Wait()
}

func Depth(B *blot.Ba) {
	//深度提取
	fmt.Println("深度提取")
}

func js_context(B *blot.Ba, url_data map[string]bool, typename string) {
	//爬取第二层应该是用一个线程，而不是一个http一个线程
	var aaa string
	for k, _ := range url_data {
		if strings.HasPrefix(k, "http") || strings.HasPrefix(k, "https") {
			//http连接
			if typename == "Ordinary" {
				continue
			} else {
				B.Get(k).Scan(&aaa)
				go js_context(B, B.Html_url(aaa), typename)
			}
		} else {
			//.js连接
			if typename == "Ordinary" {
				ordin <- k
			} else {
				dep <- k
			}
		}
	}
	if typename == "Ordinary" {
		ordin <- "True"
	} else {
		w.Add(1)
		number_de <- 1
	}
}

func go_th(B *blot.Ba) {
	//监听数据，
	for {
		select {
		case js_data := <-ordin:
			go requ_js(B, js_data, "ord")
		case js_data1 := <-dep:
			go requ_js(B, js_data1, "de")
		case number := <-number_de:
			number++
			if number == 3 {
				for i := 0; i < number; i++ {
					w.Done()
				}
			}
		}
	}

}

func requ_js(B *blot.Ba, data string, typename string) {
	l.Lock()
	defer l.Unlock()
	var js_data string
	if data == "True" {
		fmt.Println(url)
		fmt.Println(js1)
		w.Done()

	}
	js1 = append(js1, data)
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

				js1 = append(js1, data)

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

	}

}

//fuzz := config.Read_fuzz()
//for _, impression := range fuzz {
//	if ok, _ := regexp.MatchString(".*"+impression+".*", context); ok {
//		fmt.Println("包含铭感字符:", impression)
//	}
//}
