package landbalancer

import (
	"fmt"
	"testing"
)

func Test_create(t *testing.T) {
	selector := New("")
	fmt.Println(selector.Select()) // host1:port1

	selector = New("random")
	fmt.Println(selector.Select()) // host2:port2
}
