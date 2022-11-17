package config

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"os"
)

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

	file, _ := os.Open("F:\\go_project\\blot\\fuzz\\js.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//逐行读取
		line := scanner.Text() // or
		//line := scanner.Bytes()
		b = append(b, line)
		//do_your_function(line)
	}
	return

}
