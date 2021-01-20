package main

import (
	"25-build-with-meta/meta/config"
	"flag"
	"fmt"
)

var version bool

func main() {
	flag.BoolVar(&version, "version", false, "print version info")
	flag.Parse()

	if version {
		fmt.Printf("version:\t\t%s\n", config.Version)
		fmt.Printf("git hash:\t\t%s\n", config.GitHash)
		fmt.Printf("build time:\t\t%s\n", config.BuildTime)
	}
}
