package oslib

import (
	"fmt"
	"os"
	"path"
)

//OsFunctionTest 基础功能
func OsFunctionTest() {
	fmt.Println(os.Getwd())
	fmt.Println(os.Getpid())
	fmt.Println(path.Dir("/var/www/../")) // 输出 /var
}
