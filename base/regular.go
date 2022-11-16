package base

import (
	"blot/structural"
	"fmt"
	"regexp"
)

type Regular struct {
}

func (r Regular) url(resp string) (data []string) {
	//提取前端返回的url

	comp, _ := regexp.Compile(fmt.Sprintf("%s", structural.Yaml_data["regular"]))
	data = comp.FindAllString(resp, -1)
	return
}

func (r Regular) Html_url(resp string) []string {

	data := r.url(resp)
	//ii := 0
	//for index := 0; index < len(data); index++ {
	//
	//}
	fmt.Println(data)
	return data

}
