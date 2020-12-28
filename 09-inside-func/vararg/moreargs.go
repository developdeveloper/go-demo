package vararg

import "fmt"

func joinAll(prefix string, strs ...string) string {
	str := prefix

	// 可变参数的本质是切片
	fmt.Printf("%T\n", strs) // []string

	for _, s := range strs {
		str = fmt.Sprintf("%s-%s", str, s)
	}

	return str
}
