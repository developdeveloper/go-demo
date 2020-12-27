package entity

import (
	"fmt"
	"time"
)

//Book 书
type Book struct {
	Name    string
	Author  string
	Publish time.Time
}

func defineBook() {
	book := struct {
		Name    string
		Author  string
		Publish time.Time
	}{"算法导论", "大神", time.Now()}
	fmt.Printf("%T\n", book) // struct { Name string; Author string; Publish time.Time }
	fmt.Println(book)
}

func useBookStruct() {
	book1 := Book{}
	fmt.Println(book1)

	book2 := Book{
		"算法导论", "大神", time.Now(),
	}
	fmt.Println(book2)

	book3 := Book{
		Name: "算法导论", Author: "大神", Publish: time.Now(),
	}
	book3.Name = "计算机发展史"
	fmt.Println(book3)
}

func changeBookNameFailed(book Book) {
	book.Name = "无字天书"
	fmt.Printf("%p\n", &book)
	fmt.Println(book)
}

func changeBookNameSuccessed(ptrBook *Book) {
	ptrBook.Name = "无字天书"
	fmt.Printf("%p\n", ptrBook)
	fmt.Println(*ptrBook)
}

func structFuncArgs() {
	book := Book{
		"算法导论", "大神", time.Now(),
	}
	fmt.Printf("%p\n", &book)

	changeBookNameFailed(book)
	fmt.Println(book) // 算法导论

	changeBookNameSuccessed(&book)
	fmt.Println(book) // 无字天书
}
