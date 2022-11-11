package base

import (
	"regexp"
	"strings"
)

func url(resp string) (data []string) {
	//提取前端返回的url
	comp, _ := regexp.Compile("\"[^\\s\"]*/[^\\s\"]*\"")
	data = comp.FindAllString(resp, -1) //全局返回匹配的
	return
}

func Html_url(resp string) {
	//过滤掉.png jpg gif等等
	var type_data = []string{"png\"", "jpg\"", "gif\"", "css\"", "js\""}
	var data_url []string

	data := url(resp)
	//data := []string{"/css/chunk-08dac8ac.51ee81b3.css"}
	for index, i := range data {
		c := strings.Split(i, ".")
		for _, j := range type_data {
			if c[len(c)-1] == j {
				data = append(data_url, data[:index]...)
				data = append(data_url, data[:index+1]...)
				break
			}
		}
	}

}

func remove_arry() {
	//删除对应的删除
}
