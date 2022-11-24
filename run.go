package main

import (
	"fmt"
	"sync"
)

func main() {
	//a := time.Now()
	//b := blot.Start().Get("http://www.glasssix.com")
	//jsfind.Ordinary(b)
	//fmt.Println("结束")
	//fmt.Println(time.Since(a))

	go a()
	l.Add(1)
	l.Wait()
	num<-aa
	close(num)

}

var l sync.WaitGroup

var num = make(chan []int)
var aa = []int{1, 2, 3, 4, 5, 6}

func a() {
	for _, i := range aa {
		if i == 3 {
			//go a()
			fmt.Println(1)
		}
	}
	//fmt.Println("开始")
	//ll.Add(1)
	//ll.Wait()
	//
	//fmt.Println("结束")
}

var a2 int

func b() {

	for {
		select {
		case a1 := <-num:

			l.Add(1)
			a2 += 1
			fmt.Println(a1, a2)

			//fmt.Println(a2)
			if a2 == 3 {
				l.Done()
				l.Done()
				l.Done()
			}
		}
	}
}
