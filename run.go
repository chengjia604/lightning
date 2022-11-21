package main

import (
	"blot/blot"
	"blot/jsfind"
	"fmt"
)

func main() {
	//a := time.Now()
	b := blot.Start("http://www.txdyq.cn/").Get()
	jsfind.Ordinary(b)
	fmt.Println("结束")
	//fmt.Println(time.Since(a))

}
