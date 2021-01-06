package alglib

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// 按照 Person.Age 从大到小排序
type PersonSlice []*Person

func (ps PersonSlice) Len() int {
	return len(ps)
}

func (ps PersonSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps PersonSlice) Less(i, j int) bool {
	// 顺序
	return ps[j].Age > ps[i].Age

	// 逆序
	//return ps[j].Age > ps[i].Age
}

//SortPersonTest 按年龄排序
func SortPersonTest() {
	persons := PersonSlice{
		&Person{"zhangsan", 27},
		&Person{"lisi", 22},
		&Person{"wangwu", 38},
	}

	sort.Sort(persons)
	for _, person := range persons {
		fmt.Println(*person)
	}
}
