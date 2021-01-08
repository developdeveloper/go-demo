package singleton

import (
	"fmt"
	"sync"
)

type person struct{}

var instance *person
var once sync.Once

func born() *person {
	once.Do(func() {
		fmt.Println("构建 instance 对象")
		instance = &person{}
	})
	return instance
}

func getInstance() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			instance := born()
			fmt.Printf("%p\n", instance) // 0x1254790
		}()
	}

	wg.Wait()
}
