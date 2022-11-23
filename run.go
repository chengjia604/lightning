// package main
//
// import (
//
//	"blot/blot"
//	"blot/jsfind"
//	"fmt"
//	"time"
//
// )
//
//	func main() {
//		a := time.Now()
//		b := blot.Start().Get("http://www.glasssix.com")
//		jsfind.Ordinary(b)
//		fmt.Println("结束")
//		fmt.Println(time.Since(a))
//	}
package main

import (
	"fmt"
)

// write data
func writeData(intChan chan int) {
	for i := 1; i <= 55; i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeDate ", i)

	}
	close(intChan)

}

// read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到的数据 = %v\n", v)

	}

	//readData 读取完数据后，完成任务
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
