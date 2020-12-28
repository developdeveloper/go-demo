package varfunc

import (
	"fmt"
	"testing"
)

func Test_changeName(t *testing.T) {
	doubleName := changeName("zhangsan", doubleNameFunc)
	upperName := changeName("lisi", upperNameFunc)
	fmt.Println(doubleName) // zhangsanzhangsan
	fmt.Println(upperName)  // LISI
}
