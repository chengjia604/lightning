package blot

import (
	"blot/structural"
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
)

type Regular struct {
}

func (r Regular) url(resp string) (data []string) {

	//comp, _ := regexp.Compile(fmt.Sprintf("%s", structural.Yaml_data["regular"]))
	//data = comp.FindAllString(resp, -1)
	//return
	re := regexp2.MustCompile(fmt.Sprintf("%s", structural.Yaml_data["regular"]), 0)
	m, _ := re.FindStringMatch(resp)
	for m != nil {
		data = append(data, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return
}

var data_map = make(map[string]bool)

func (r Regular) Html_url(resp string) map[string]bool {
	/*数据清洗*/
	data := r.url(resp)
	mismatch := []string{"jpg", "png", "gif", "jpeg", "css", "ico"}
	for _, i := range data {
		c := strings.Split(i, "\"")
		c1 := strings.Split(c[len(c)-1], ".")
		a := c1[len(c1)-1] //取后缀
		for _, houzui := range mismatch {
			if a == houzui {
				goto c
			}
		}
		data_map[c[len(c)-1]] = true
	c:
		continue

	}

	//ii := 0
	//num_index := len(data)
	//for index := 0; index < num_index; index++ {
	//	var b = true
	//	data_index := data[index-ii]
	//	c := strings.Split(data_index, "\"")
	//	c1 := strings.Split(c[len(c)-1], ".")
	//	a := c1[len(c1)-1] //取后缀
	//	for _, index_h := range mismatch {
	//		//判断后缀
	//		if a == index_h {
	//			data = append(data[:index-ii], data[index+1-ii:]...) //过滤
	//			ii++
	//			b = false
	//			break
	//
	//		}
	//	}
	//	if b {
	//		data[index-ii] = c[len(c)-1]
	//	}
	//
	//}

	return data_map
}

func (r Regular) Domain(url string) string {
	//带域名提取
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

func (r Regular) Subdomain(url string) (string, string) {
	//不带http域名，第一个返回是xxx.xxx.com域名，第二个是xxx.com

	re := regexp2.MustCompile(fmt.Sprintf("%s", structural.Yaml_data["subdomain"]), 0)
	//if url == structural.Yaml_data["domain"] {
	//	return ""
	//}

	subdomain, _ := re.FindStringMatch(url)
	if subdomain != nil {
		a := strings.Split(subdomain.String(), ".")
		var b string
		if len(a) >= 3 {
			b = a[1] + "." + a[2]
			return subdomain.String(), b
		} else {
			return subdomain.String(), subdomain.String()
		}
	}
	panic(err)

}
