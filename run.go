package main

import (
	"fmt"
	"strings"
)

func main() {
	//var a string
	//b := blot.Start()
	//b.Get("https://home.firefoxchina.cn/?from=extra_start").Scan(&a)
	//b.Html_url(a)

	var a = "src=\"123.png"
	fmt.Println(strings.Split(a, "(src|herf)=\""))
}
