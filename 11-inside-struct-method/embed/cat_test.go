package embed

import (
	"fmt"
	"testing"
)

func Test_cat(t *testing.T) {
	where := addr{address: "东大街", phone: "123"}
	c := cat{color: "黑红", addr: where}
	fmt.Println(c)
}

func Test_dog(t *testing.T) {
	p := person{sex: "女", addr: addr{address: "南大街", phone: "321"}}
	fmt.Println(p)
}
