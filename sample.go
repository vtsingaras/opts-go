package main

import (
	"fmt"
	"./opts"
)

func printVersion() {
	fmt.Printf("VERSION\n")
}

var output = opts.Half("-o", "", "write output to file", "", "output")
var showVersion = opts.Flag("", "--version", "Description")

func main() {
	opts.Parse()
	if *showVersion {
		printVersion()
	}
	if *output != "" {
		fmt.Printf("Writing output to %s\n", *output)
	}
}
