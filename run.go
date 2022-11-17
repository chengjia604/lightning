package main

import (
	"blot/config"
	"fmt"
	"regexp"
)

func main() {
	//b := blot.Start("https://www.xiuzhanwang.com/a1").Get()
	//jsfind.Ordinary(b)
	for _, impression := range config.Read_fuzz() {
		if ok, _ := regexp.MatchString(".*"+impression+".*", "js_dataser123"); ok {
			fmt.Println(1)
		}
	}
}
