package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_person(t *testing.T) {
	var what interface{}
	what = person{"zhangsan", 20}

	// fmt.Println(reflect.TypeOf(what))
	// fmt.Println(reflect.ValueOf(what))

	val := reflect.TypeOf(what)
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := reflect.ValueOf(val.Field(i))
			fmt.Println(field)
		}
	default:
		fmt.Println("not struct")
	}
}
