package main

import (
	"fmt"
	"reflect"
)

func caseType(what interface{}) {
	switch what.(type) {
	case int:
		what = what.(int) + 1
		fmt.Printf("number %d\n", what)
	case string:
		fmt.Printf("string %s\n", what.(string)+" suffix")
	case nil:
		fmt.Println("you are nothing")
	}
}

func detectType(what interface{}) {
	val := reflect.ValueOf(what)
	//t := reflect.TypeOf(val)
	t := val.Type()       // Type 是接口对象
	fmt.Printf("%T\n", t) // *reflect.rtype 表示通用的类型结构
	//fmt.Println(t.Name()) // 类型的名称，比如 reflect.Int 打印 int

	if val.Kind() == reflect.Int {
		number := val.Int() + 1
		fmt.Printf("number %d\n", number) // number 11
	} else if val.Kind() == reflect.String {
		fmt.Printf("string %s\n", val.String()+" suffix") // string test suffix
	}
}

func getReflectValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	// Uintptr
	// UnsafePointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func reflectIterate(x interface{}, handler func(val int)) {
	val := getReflectValue(x)

	switch val.Kind() {
	case reflect.Int, reflect.Uint, reflect.Float32, reflect.Complex64:
		// 其他不同位数的数值类型
		// 以简单值调用 handler
	case reflect.Bool:
		// 以简单值调用 handler
	case reflect.String:
		// 以简单值调用 handler
	case reflect.Interface:
		// 特殊的
	case reflect.Chan:
		// 特殊的
	case reflect.Func:
		// 特殊的
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			reflectIterate(val.Field(i).Interface(), handler)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			reflectIterate(val.Index(i).Interface(), handler)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			reflectIterate(val.MapIndex(key).Interface(), handler)
		}
	}
}
