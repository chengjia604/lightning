package main

import (
	"blot/blot"
	"fmt"
	"time"
)

func main() {

	var a string
	b := blot.Start()
	bT := time.Now()
	b.Get("https://www.xiuzhanwang.com/a1/").Scan(&a)
	eT := time.Since(bT)
	fmt.Println(eT)
	b.Html_url(a)

	// 开始时间

}
