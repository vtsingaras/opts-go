package main

import (
	"opts"
	"fmt"
)

func main() {
	opts.Shortflag("h", "get help")
	opts.Parse()
	fmt.Printf("Xname: %s\n", opts.Xname)
	for i := 0; i < len(opts.Args); i++ {
		fmt.Printf("Argument: %s\n", opts.Args[i])
	}
}
