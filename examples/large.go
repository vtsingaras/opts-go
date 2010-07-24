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
	opts.Description("a sample program")
	opts.Parse()
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
