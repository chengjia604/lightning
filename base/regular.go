package base

import (
	"fmt"
	"regexp"
	"strings"
)

const a = "\"[^\\s]*/[^\\s]*\""

func url(resp string) (data []string) {
	//提取前端返回的url

	comp, _ := regexp.Compile(a)
	data = comp.FindAllString(resp, -1) //全局返回匹配的

	return
}

func Html_url(resp string) {
	//过滤掉.png jpg gif等等
	var type_data = []string{"png\"", "jpg\"", "gif\"", "css\"", "js\""}
	//var url_data []string
	data := url(resp)
	ii := 0
	//index := 0

	for index := 0; index < len(data); index++ {
		c := strings.Split(data[index], ".")
		for _, j := range type_data {
			if c[len(c)-1] == j {
				data = append(data[:index-ii], data[index+1-ii:]...)
				ii++
				break
			}
		}

	}
	fmt.Println(data)
	//for _, i := range data {
	//	c := strings.Split(i, ".")
	//	for _, j := range type_data {
	//		if c[len(c)-1] == j {
	//			data = append(data[:index-ii], data[index-ii+1:]...)
	//			ii++
	//			break
	//		}
	//	}
	//	index++
	//}

}

func remove_arry() {
	//删除对应的删除
}
