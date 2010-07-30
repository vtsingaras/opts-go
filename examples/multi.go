package main

import (
	"fmt"
	"opts"
)

func main() {
	include := opts.Multi("I", "include", "add include path")
	opts.Description("a sample program")
	opts.Parse()
	for _, path := range *include {
		fmt.Printf("Adding to include path: %s\n", path)
	}
}
