package keyvalue

import (
	"fmt"
	"testing"
)

func Test_updateDict(t *testing.T) {
	dict := map[string]string{"one": "1", "two": "2", "three": "3"}
	updateDict(dict)
	fmt.Println(dict) // map[four:4 three:3 two:2]
}
