package structural

import (
	"blot/config"
	"fmt"
)

type yaml_data struct {
	data map[any]any
}

var Yaml_data = yaml_data{config.Read_config()}.data
var Useraget = fmt.Sprintf("%s", Yaml_data["headers"].(map[any]any)["user-agent"])
