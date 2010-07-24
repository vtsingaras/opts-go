package main

import (
	"fmt"
	"opts"
	"os"
)

func main() {
	help := opts.Flag("h", "help", "get help")
	v := opts.Longflag("version", "print version information")
	world := opts.Shortopt("v", "the string to use instead of 'world'", "world")
	file := opts.Option("f", "file", "a file to be looked at", "myfile")
	quiet := opts.Shortflag("q", "be quiet")
	opts.Parse()
	fmt.Printf("Xname: %s\n", opts.Xname)
	for i := 0; i < len(opts.Args); i++ {
		fmt.Printf("Argument: %s\n", opts.Args[i])
	}
	if *help {
		opts.Help()
		os.Exit(0)
	}
	if *v {
		fmt.Printf("Hello, world! for opts.go\n")
		os.Exit(0)
	}
	if *quiet {
		fmt.Printf("I'm being quiet!\n")
	}
	fmt.Printf("Hello, %s\n", *world)
	fmt.Printf("Reading stuff from %s\n", *file)
}
