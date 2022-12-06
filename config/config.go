package config

import (
	"bufio"
	"encoding/json"
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

func Create_html(url, domname map[string]bool, name string) {
	//创建模板
	file, err := os.OpenFile(filepath.Join(rpath(), "tem/"+name+".html"), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic("生成模板错误")
	}
	defer file.Close()
	file_data := bufio.NewWriter(file)
	file_data.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"utf-8\">\n<link rel=\"stylesheet\" type=\"text/css\" href='./index.css'>\n    <script src=\"https://cdn.bootcss.com/jquery/3.3.1/jquery.js\"></script>\n    <script src=\"https://cdn.bootcss.com/jquery.form/4.2.2/jquery.form.js\"></script>\n<title>Lightning</title>\n</head>\n<body>\n<div class=\"sidenav\">\n    <a href=\"\">url资产</a>\n    <a href=\"\" id=\"domname\">其他域名资产</a>\n    <a href=\"\" id=\"js\">js资产</a>\n</div>\n    <div>\n        <table border=\"1\" cellspacing=\"0\" class=\"table\" id=\"tb\">\n            <tr>\n                <th>url资产</th>\n                <th>敏感字段</th>\n            </tr>\n       </table>\n    </div>\n \n</body>\n<script src=\"./index.js\"></script>    \n</html>")
	//for k, _ := range url {
	//	a := fmt.Sprintf("<tr>\n        <td>%s</td>\n        <td></td>\n    </tr>", k)
	//	file_data.WriteString(a)
	//}
	//file_data.WriteString("</table>\n \n</body>\n</html>")
	file_data.Flush()
	js, err1 := os.OpenFile(filepath.Join(rpath(), "tem/index.js"), os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		panic("生成js文件错误")
	}
	json_url, _ := json.Marshal(url)
	json_domname, _ := json.Marshal(domname)
	js_data := bufio.NewWriter(js)
	js_data.WriteString(fmt.Sprintf("var url=[JSON.parse('%s'.replace(/\\//g, ''))];\nvar domnameurl=[JSON.parse('%s'.replace(/\\//g, ''))];\nvar jsname=[JSON.parse('%s'.replace(/\\//g, ''))];\n", json_url, json_domname, "1"))
	js_data.WriteString("var domname=document.getElementById(\"domname\");\ndomname.onclick=function(){\n    for(let [key,value] of Object.entries(domnameurl[0])){\n        let html= `\n        <tr>\n        <td>${key}</td>\n         <td>${value}</td>\n        </tr>\n        `\n        $('#tb tbody').append(html);\n    }\n}\nfor(let [key,value] of Object.entries(url[0])){\n        let html= `\n        <tr>\n        <td>${key}</td>\n         <td>${value}</td>\n        </tr>\n        `\n        $('#tb tbody').append(html);\n}")
	js_data.Flush()
}
