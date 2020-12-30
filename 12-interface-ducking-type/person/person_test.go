package person

import "testing"

func Test_person(t *testing.T) {
	p := Person{"zhangsan", 25, 20, 10}
	p.Walk(100)
	p.Swim(100)
}
