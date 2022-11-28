package config

import (
	"bufio"
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
