package assert

import (
	"testing"
)

func Test_person(t *testing.T) {
	var what interface{}
	what = &person{"zhangsan", 20}

	if p, ok := what.(*person); ok {
		p.eat()
	}
}
