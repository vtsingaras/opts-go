package main

import (
	"opts"
	"fmt"
)

func main() {
	opts.Shortopt("h", "get help", "false")
	opts.Parse()
	fmt.Printf("Xname: %s\n", opts.Xname)
}
