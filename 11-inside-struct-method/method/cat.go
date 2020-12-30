package method

import "fmt"

type cat struct {
	name  string
	color string
	age   int
}

func (c cat) run() {
	fmt.Println("running")
}

func (c cat) jump() {
	fmt.Println("jumping")
}

func (c cat) grow1() {
	fmt.Printf("%p\n", &c) // 0xc000092390
	c.age++
}

func (c *cat) grow2() {
	fmt.Printf("%p\n", c) // 0xc000092360
	c.age++
}
