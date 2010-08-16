package main

import (
	"fmt"
	"./opts"
)

func printVersion() {
	fmt.Printf("VERSION\n")
}

var showVersion = opts.Flag("", "--version", "Description")

func main() {
	opts.Parse()
	if *showVersion {
		printVersion()
	}
}
