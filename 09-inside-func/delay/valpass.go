package delay

import "fmt"

func readName(name string) {
	fmt.Println("changeName " + name) // changeName zhangsan
}

func passName() {
	name := "zhangsan"
	defer readName(name)
	name = "lisi"
	fmt.Println("passName " + name) // passName lisi
}

func readNumberRef(ptrNumber *int) {
	fmt.Printf("readNumberRef %d\n", *ptrNumber) // readNumberRef 1
}

func passNumberRef() {
	number := 0
	defer readNumberRef(&number)
	number = 1
	fmt.Printf("passNumberRef %d\n", number) // passNumberRef 1
}

func wontRun() {
	if false {
		defer fmt.Println("should not see me") // 不会执行
	}

	fmt.Println("see me")
}

type book struct {
	Name string
}

func printBook(b *book) {
	fmt.Println(*b)
}

func walkBooks() {
	books := []book{
		{Name: "计算机"}, {Name: "电脑"},
	}
	// for _, b := range books {
	// 	defer printBook(&b)
	// }

	for index := range books {
		defer printBook(&books[index])
	}
}

