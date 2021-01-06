package iolib

import (
	"bufio"
	"fmt"
	"strings"
)

//CountWordsTest 单词计数
func CountWordsTest(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	total := 0

	// 返回 false 时停止扫描，可能扫描完了，也可能是出错了
	// 需要判断 Err() 函数的值来确定
	for scanner.Scan() {
		total++
	}

	// 不会是 io.EOF 错误
	if err := scanner.Err(); err != nil {
		fmt.Println("出错了:" + err.Error())
	}

	fmt.Printf("包含 %d 个单词\n", total)
}
