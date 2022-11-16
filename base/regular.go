package base

import (
	"blot/structural"
	"fmt"
	"regexp"
)

type Regular struct {
}

func (r Regular) url(resp string) (data []string) {

	comp, _ := regexp.Compile(fmt.Sprintf("%s", structural.Yaml_data["regular"]))
	data = comp.FindAllString(resp, -1)
	return
}

func (r Regular) Html_url(resp string) []string {
	/*数据清洗*/
	data := r.url(resp)
	ii := 0
	re, _ := regexp.Compile("\"")
	for index := 0; index < len(data); index++ {
		if ok, _ := regexp.MatchString("(src|herf)=*", data[index-ii]); ok {
			re.FindString()
		} else {
			continue
		}
	}
	fmt.Println(data)
	return data

}
