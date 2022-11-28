package main

import (
	"blot/blot"
	"blot/jsfind"
	"fmt"
	"time"
)

func main() {
	//host := flag.String("host", "127.0.0.1", "请输入host地址")
	//port := flag.Int("port", 3306, "请输入端口号")
	//flag.Parse() // 解析参数
	//fmt.Printf("%s:%d\n", *host, *port)
	//fmt.Println("==================================================================================================================================================================================")
	//color.Blue(" .----------------.  .----------------.  .----------------.  .----------------. \n| .--------------. || .--------------. || .--------------. || .--------------. |\n| |   ______     | || |     ____     | || |   _____      | || |  _________   | |\n| |  |_   _ \\    | || |   .'    `.   | || |  |_   _|     | || | |  _   _  |  | |\n| |    | |_) |   | || |  /  .--.  \\  | || |    | |       | || | |_/ | | \\_|  | |\n| |    |  __'.   | || |  | |    | |  | || |    | |   _   | || |     | |      | |\n| |   _| |__) |  | || |  \\  `--'  /  | || |   _| |__/ |  | || |    _| |_     | |\n| |  |_______/   | || |   `.____.'   | || |  |________|  | || |   |_____|    | |\n| |              | || |              | || |              | || |              | |\n| '--------------' || '--------------' || '--------------' || '--------------' |\n '----------------'  '----------------'  '----------------'  '----------------'")
	//flag.StringVar(&U, "u", "", "获取的域名")
	//flag.Parse()
	//start()

	a := time.Now()
	b := blot.Start().Get("https://www.glasssix.com/")
	jsfind.Ordinary(b)
	fmt.Println(time.Since(a))
}

var U string

func start() {
	a := time.Now()
	b := blot.Start().Get(fmt.Sprintf("%v", U))
	jsfind.Ordinary(b)
	fmt.Println("所用耗时：", time.Since(a))
}
