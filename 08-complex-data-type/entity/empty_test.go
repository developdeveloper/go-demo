package entity

import (
	"fmt"
	"testing"
)

func Test_lookupEmptyStruct(t *testing.T) {
	lookupEmptyStruct()
}

func Test_set(t *testing.T) {
	// type Set map[string]struct{}{}
	set := map[string]struct{}{}
	set["one"] = struct{}{}
	set["two"] = struct{}{}
	set["three"] = struct{}{}
	fmt.Println(set["notfound"])
}
