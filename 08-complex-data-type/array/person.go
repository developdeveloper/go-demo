package array

import "fmt"

//callcppclass 人
type person struct {
	Name string
	Age  int
}

func forRangeStruct() {
	persons := []person{
		{"zhangsan", 20},
		{"lisi", 21},
	}

	for index, item := range persons {
		fmt.Printf("%p\n", &item)           // 0xc00012a0a0
		fmt.Printf("%p\n", &persons[index]) // 0xc000108360, 0xc000108378
		fmt.Println(index, item)
	}
}
