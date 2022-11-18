package blot

import (
	"blot/structural"
	"fmt"
	"regexp"
	"strings"
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
	mismatch := []string{"jpg", "png", "gif", "jpeg"}
	ii := 0
	num_index := len(data)
	for index := 0; index < num_index; index++ {
		var b = true
		data_index := data[index-ii]
		if ok, _ := regexp.MatchString("(src|herf)=*", data_index); ok {
			c := strings.Split(data_index, "\"")
			c1 := strings.Split(c[len(c)-1], ".")
			a := c1[len(c1)-1]
			for _, index_h := range mismatch {
				//判断后缀
				if a == index_h {
					data = append(data[:index-ii], data[index+1-ii:]...)
					ii++
					b = false
					break
				}
			}
			if b {
				data[index-ii] = c[len(c)-1]
			}
		} else {
			continue
		}
	}
	return data
}

func (r Regular) Domain(url string) (domain string) {
	//域名提取
	re, _ := regexp.Compile(fmt.Sprintf("%s", structural.Yaml_data["domain"]))
	domain = re.FindString(url)
	return
}
