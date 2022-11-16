package main

import "blot/blot"

func main() {
	var a string
	b := blot.Start()
	b.Get("https://www.xiuzhanwang.com/a1/").Scan(&a)
	b.Html_url(a)

}
