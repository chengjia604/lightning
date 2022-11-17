package main

import (
	"blot/blot"
	"blot/jsfind"
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	b := blot.Start("http://ehome.homekoo.com/login.php").Get()
	jsfind.Ordinary(b)
	fmt.Println("结束")
	fmt.Println(time.Since(a))
}
