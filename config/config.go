package config

import (
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
