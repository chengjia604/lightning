package structural

import "blot/config"

type yaml_data struct {
	data map[any]any
}

var Yaml_data = yaml_data{config.Read_config()}.data
