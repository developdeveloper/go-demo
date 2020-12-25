package main

import "fmt"

func main() {
	doNothing()

	// 直接编码需要实现功能的细节
	// fmt.Println("boring\thello")
	// fmt.Println("boring\thello")

	// 用函数封装后重用
	// 不在乎怎么说 "hello" 的功能，能完成说  "hello" 的功能就行了
	sayHello()
	sayHello()

	holdFuncWithVar()

	fmt.Println(incNumber(1))

	fmt.Println(swapStr("1st", "2nd"))

	fmt.Println(sumAllWithInitial(0, 1, 2, 3))

	builtInFunc()
}

func doNothing() {
	// 函数体
}

func sayHello() {
	fmt.Println("hello")
}

// 直接调用一个匿名函数
func anonymous() {
	func() {
		fmt.Println("i am anonymous function")
	}() // 注意这个括号
}

// 把函数复制给一个变量
func holdFuncWithVar() {
	f := func(name string) {
		fmt.Printf("hey, %s\n", name)
	}

	// 显示 func(string)
	// fmt.Printf("%T", f)

	f("zhangsan")
}

// 函数有返回值
func incNumber(in int) int {
	return in + 1
}

// 函数有多个返回值
func swapStr(s1, s2 string) (string, string) {
	return s2, s1
}

// more 是不定参数，result 是返回变量
func sumAllWithInitial(initial int, more ...int) (result int) {
	result = initial

	// 循环读取 more 参数，以后会学习
	for _, num := range more {
		result = result + num
	}

	// 等效于 return result
	return
}

func builtInFunc() {
	str := "english中文"
	// len 是一个内置函数
	// 对于 string 它返回字节
	fmt.Println(len(str)) // 返回 13，因为 1 个中文 UTF-8 这里存储占 3 个字节
}
