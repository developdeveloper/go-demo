package file

import (
	"fmt"
	"os"
)

func read() {
	file, err := os.Open("/tmp/test")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	file.Close()
}
