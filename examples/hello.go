package main

import (
	"opts"
	"fmt"
)

func main() {
	help := opts.Flag("h", "help", "get help")
	v := opts.Longflag("version", "print version information") 
	opts.Parse()
	fmt.Printf("Xname: %s\n", opts.Xname)
	for i := 0; i < len(opts.Args); i++ {
		fmt.Printf("Argument: %s\n", opts.Args[i])
	}
	if *help {
		fmt.Printf("help stuff here\n")
	}
	if *v {
		fmt.Printf("version information here\n")
	}
}
