package main

import (
	"fmt"
	"regexp"
)

func main() {
	//a := time.Now()
	//b := blot.Start("http://ehome.homekoo.com/login.php").Get()
	//jsfind.Ordinary(b)
	//fmt.Println("结束")
	//fmt.Println(time.Since(a))
	if ok, _ := regexp.MatchString("(https|http?://)[^/]+([[^:blank/]]*)", "https://ehome.homekoo.com/login.php"); ok {
		fmt.Println(1)
	}
}
