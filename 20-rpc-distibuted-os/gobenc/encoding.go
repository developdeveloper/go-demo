package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"reflect"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{"zhangsan", 20}
	enc1 := gob.NewEncoder(os.Stdout)
	enc1.Encode(p1)

	file1, _ := os.OpenFile("/tmp/callcppclass.gob", os.O_CREATE|os.O_WRONLY, 0644)
	enc2 := gob.NewEncoder(file1)
	enc2.Encode(p1)
	file1.Close()

	var p2 person
	file2, _ := os.Open("/tmp/callcppclass.gob")
	defer file2.Close()
	dec := gob.NewDecoder(file2)
	dec.Decode(&p2)

	fmt.Println(p2)
	if reflect.DeepEqual(p1, p2) {
		fmt.Println("decode ok")
	}
}
