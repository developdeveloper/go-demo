package main

import (
	"fmt"
	// "unsafe"
)

func main() {
	printHello()
	printSpace()
	printWorld()

	printChinese()
	numberOverflow()
	printBool()
	printInitalText()
}

func printSpace() {
	// var space uint8
	var space byte
	space = 32
	// space = '\x20'
	fmt.Printf("%c", space)
	// fmt.Printf("%T\n", space)
	// fmt.Printf("%d", unsafe.Sizeof(space))
}

func printHello() {
	str := "hello"
	// fmt.Printf("%T\n", str)
	fmt.Printf(str)
}

func printWorld() {
	var str string
	str = "world"
	// var str string = "world"
	fmt.Printf(str)
}

func printChinese() {
	var ch1 rune = '中'
	var ch2 rune = '国'
	fmt.Printf("\n%c%c\n", ch1, ch2)
	// fmt.Printf("%d\n", unsafe.Sizeof(ch1))
	// fmt.Printf("%d\n", unsafe.Sizeof(ch2))
}

func numberOverflow() {
	var num1 uint8 = 50 // 最大表是 256
	var num2 uint8 = 50 // 最大表是 256
	var num3 uint8

	num3 = num1 * num2
	fmt.Printf("%d", num3) // 期望是 2500，结果显示 196，改为 uint16 则正常
}

func printBool() {
	b := true
	fmt.Printf("\n%T", b)
}

func printInitalText() {
	str := `I hold keep
			myself,
						\r\n\t\b
can you?
	`
	fmt.Println()
	fmt.Printf(str)
}
