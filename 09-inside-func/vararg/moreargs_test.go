package vararg

import (
	"fmt"
	"testing"
)

func Test_joinAll(t *testing.T) {
	str := joinAll("prefix", "one", "two", "three")
	fmt.Println(str) // prefix-one-two-three
}
