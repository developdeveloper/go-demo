package main

import (
	"24-parse-cmd-flags/slicearg"
	"flag"
	"fmt"
	"os"
)

var (
	help   bool
	color  bool
	hours  int
	fruits []string
)

func main() {
	//fmt.Println(os.Args)

	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&color, "color", false, "apply color display")
	flag.IntVar(&hours, "hours", 24, "date time mode")

	flag.Var(slicearg.New(&fruits, []string{"nothing"}), "fruits", "do you like sth")

	flag.Usage = usage
	flag.Parse()

	if help {
		return
	}

	if hours == 24 {
		fmt.Println("23:59:59")
	} else {
		fmt.Println("not 24 hours mode")
	}

	fmt.Println(fruits)
}

func usage() {
	fmt.Fprintf(os.Stderr, "你好，这是自定义帮助信息\n")
	flag.PrintDefaults()
}
