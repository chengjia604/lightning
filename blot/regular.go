package blot

import (
	"blot/structural"
	"fmt"
	"github.com/dlclark/regexp2"
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
	mismatch := []string{"jpg", "png", "gif", "jpeg", "css", "ico"}
	ii := 0
	num_index := len(data)
	for index := 0; index < num_index; index++ {
		var b = true
		data_index := data[index-ii]
		c := strings.Split(data_index, "\"")
		c1 := strings.Split(c[len(c)-1], ".")
		a := c1[len(c1)-1] //取后缀
		for _, index_h := range mismatch {
			//判断后缀
			if a == index_h {
				data = append(data[:index-ii], data[index+1-ii:]...) //过滤
				ii++
				b = false
				break
			}
		}
		if b {
			data[index-ii] = c[len(c)-1]
		}
	}
	return data
}

func (r Regular) Domain(url string) string {
	//域名提取

	re := regexp2.MustCompile(fmt.Sprintf("%s", structural.Yaml_data["domain"]), 0)
	//if url == structural.Yaml_data["domain"] {
	//	return ""
	//}

	domain, _ := re.FindStringMatch(url)
	if domain != nil {

		return domain.String()
	}
	return url

}
