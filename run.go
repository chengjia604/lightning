package main

import (
	"blot/blot"
	"blot/jsfind"
	"blot/structural"
	"flag"
	"fmt"
	"github.com/fatih/color"
	_ "net/http/pprof"
	"time"
)

var u string

func init() {
	flag.StringVar(&u, "u", "https://www.glasssix.com/", "获取的域名")
	flag.StringVar(&blot.Cookie, "cookie", "", "设置cookie")
	flag.StringVar(&structural.Useraget, "us", "", "设置useraget")
	flag.StringVar(&blot.I, "i", "456", "生成文档")
	flag.StringVar(&blot.S, "s", "", "详细显示")

}

func main() {
	//color.Blue(" .----------------.  .----------------.  .----------------.  .----------------. \n| .--------------. || .--------------. || .--------------. || .--------------. |\n| |   ______     | || |     ____     | || |   _____      | || |  _________   | |\n| |  |_   _ \\    | || |   .'    `.   | || |  |_   _|     | || | |  _   _  |  | |\n| |    | |_) |   | || |  /  .--.  \\  | || |    | |       | || | |_/ | | \\_|  | |\n| |    |  __'.   | || |  | |    | |  | || |    | |   _   | || |     | |      | |\n| |   _| |__) |  | || |  \\  `--'  /  | || |   _| |__/ |  | || |    _| |_     | |\n| |  |_______/   | || |   `.____.'   | || |  |________|  | || |   |_____|    | |\n| |              | || |              | || |              | || |              | |\n| '--------------' || '--------------' || '--------------' || '--------------' |\n '----------------'  '----------------'  '----------------'  '----------------'")
	color.Blue("  _ _       _     _         _             \n | (_)     | |   | |       (_)            \n | |_  __ _| |__ | |_ _ __  _ _ __   __ _ \n | | |/ _` | '_ \\| __| '_ \\| | '_ \\ / _` |\n | | | (_| | | | | |_| | | | | | | | (_| |\n |_|_|\\__, |_| |_|\\__|_| |_|_|_| |_|\\__, |\n       __/ |                         __/ |\n      |___/                         |___/ ")
	color.Yellow("名称：闪电")
	color.Yellow("版本: 1.0.1")
	color.Yellow("语言: Go")
	color.Yellow("=================================")
	flag.Parse()

	if u == "" {
		color.Green("似乎没有域名")
		return
	}
	color.Green("数据收集中...." + "\n")
	start()
}

func start() {
	a := time.Now()
	b := blot.Start().Get(fmt.Sprintf("%v", u))
	//jsfind.Ordinary(b)
	jsfind.Ord(b)
	fmt.Println("所用耗时:", time.Since(a))
}
