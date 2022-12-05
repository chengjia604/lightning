package config

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func rpath() string {
	project_path, _ := os.Getwd()
	return project_path
}

func Read_config() (map_data map[any]any) {
	/*读取配置文件*/
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		panic("读取配置文件错误")
	}
	if err1 := yaml.Unmarshal(data, &map_data); err != nil {
		panic(err1)
	}
	return
}

func Read_fuzz() (b []string) {
	//读取fuzz

	file, err := os.Open(filepath.Join(rpath(), "fuzz/js.txt"))
	if err != nil {
		panic("读取fuzz错误")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//逐行读取
		line := scanner.Text() // or
		//line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		b = append(b, line)
		//do_your_function(line)
	}
	return
}

func Create_html(url map[string]bool) {
	//创建模板

	file, err := os.OpenFile(filepath.Join(rpath(), "tem/index.html"), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic("生成模板错误")
	}
	defer file.Close()
	file_data := bufio.NewWriter(file)
	file_data.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"utf-8\">\n<link rel=\"stylesheet\" type=\"text/css\" href='./index.css'>\n<title>Lightning</title>\n</head>\n<body>\n<div class=\"sidenav\">\n    <a href=\"/url资产\">url资产</a>\n    <a href=\"/其他域名资产\">其他域名资产</a>\n    <a href=\"/js资产\">js资产</a>\n</div>\n    <div>\n        <table border=\"1\" cellspacing=\"0\" class=\"table\">\n            <tr>\n                <th>url资产</th>\n                <th>敏感字段</th>\n            </tr>")
	for k, _ := range url {
		a := fmt.Sprintf("<tr>\n        <td>%s</td>\n        <td></td>\n    </tr>", k)
		file_data.WriteString(a)
	}
	file_data.WriteString("</table>\n \n</body>\n</html>")
	file_data.Flush()
}
