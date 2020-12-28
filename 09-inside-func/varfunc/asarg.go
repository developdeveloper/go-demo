package varfunc

import "strings"

// NameExchangeFunc 变换名称的函数
type NameExchangeFunc func(string) string

func doubleNameFunc(name string) string {
	return name + name
}

func upperNameFunc(name string) string {
	return strings.ToUpper(name)
}

// func changeName(name string, f func(string) string) string {
func changeName(name string, f NameExchangeFunc) string {
	return f(name)
}
