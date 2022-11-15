package base

import (
	"fmt"
	"regexp"
)

func url(resp string) (data []string) {
	//提取前端返回的url

	comp, _ := regexp.Compile("")
	data = comp.FindAllString(resp, -1) //全局返回匹配的
	return
}

func Html_url(resp string) {

	data := url(resp)
	fmt.Println(data)
	//ii := 0
	//for index := 0; index < len(data); index++ {
	//	c := strings.Split(data[index-ii], ".")
	//	for _, j := range type_data {
	//		if c[len(c)-1] == j {
	//			data = append(data[:index-ii], data[index+1-ii:]...)
	//			ii++
	//			break
	//		}
	//	}
	//
	//}
	//fmt.Println(data)
}

func remove_arry() {
	//删除对应的删除
}
