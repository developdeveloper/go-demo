package method

import (
	"fmt"
	"testing"
)

func Test_cat(t *testing.T) {
	myCat := cat{"虎斑猫", "黑黄", 1}
	fmt.Printf("%p\n", &myCat) // 0xc000092360

	myCat.run()
	myCat.jump()
	myCat.grow1()
	fmt.Println(myCat) // {虎斑猫 黑黄 1}

	ptrMyCat := &myCat
	ptrMyCat.run()
	ptrMyCat.jump()
	ptrMyCat.grow2()
	fmt.Println(*ptrMyCat) // {虎斑猫 黑黄 2}
}
