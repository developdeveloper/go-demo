package assert

import "fmt"

type person struct {
	name string
	age  uint8
}

func (p person) eat() {
	fmt.Println("eating")
}
