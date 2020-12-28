package retval

import (
	"fmt"
	"testing"
)

func Test_getNameFunc(t *testing.T) {
	f1 := getNameFunc("zhangsan")
	f2 := getNameFunc("lisi")
	fmt.Println(f1()) // zhangsan
	fmt.Println(f2()) // lisi
}
