package delay

import (
	"fmt"
	"os"
)

func readfile() string {
	file, _ := os.Open("/tmp/command.txt") // 示例忽略错误
	buf := make([]byte, 10)
	size, _ := file.Read(buf)
	str := string(buf[0:size])
	fmt.Println(str)

	if str == "start" {
		file.Close()
		return "execute service start"
	} else if str == "stop" {
		file.Close()
		return "execute service stop"
	}

	file.Close()
	return "unknow"
}

func readfile2() string {
	file, _ := os.Open("/tmp/command.txt") // 示例忽略错误
	defer file.Close()

	buf := make([]byte, 10)
	size, _ := file.Read(buf)
	str := string(buf[0:size])
	fmt.Println(str)

	if str == "start" {
		return "execute service start"
	} else if str == "stop" {
		return "execute service stop"
	}

	return "unknow"
}
