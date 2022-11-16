package main

import (
	"fmt"
	"regexp"
)

func main() {
	//var a string
	//b := blot.Start()
	//b.Get("https://home.firefoxchina.cn/?from=extra_start").Scan(&a)
	//b.Html_url(a)
	re, _ := regexp.Compile("\".*")
	var a = "src=\"123.png"
	fmt.Println(re.FindString(a))
}
