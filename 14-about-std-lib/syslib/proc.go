package syslib

import (
	"fmt"
	"syscall"
)

//GetCurrentPid 当前进程ID
func GetCurrentPidTest() {
	fmt.Println(syscall.Getpid())
}
