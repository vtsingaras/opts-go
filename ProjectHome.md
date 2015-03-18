opts.go is lightweight package for POSIX- and GNU-style options for go.

Important features that opts.go has:

  * Automatic help screen generation
  * Support for both short and long options (`-o` and `--opt`)
  * Support for multiple short options in one argument
    * Including the case where the last option requires an argument
  * Automatically allows both `-I file` and `-Ifile` formats.
  * Support for options being specified multiple times, with different values
  * Support for optional arguments.

Things opts.go lacks, by design:

  * All arguments to options are strings.
  * opts.go does not support multiple mandatory arguments to a single option

```
package main

import (
 "fmt"
 "./opts"
)

func printVersion() {
 fmt.Printf("VERSION\n")
}

var showVersion = opts.Flag("", "--version", "Description")
var output = opts.Half("-o", "--output", "write output to file", "", "output")
var format = opts.Single("-f", "--format", "the output format to use", "csv")
var include = opts.Multi("-I", "--include", "files to include", "file")

func main() {
        opts.Parse()
        if *showVersion {
                printVersion()
        }
        if *output != "" {
                fmt.Printf("Writing output to %s\n", *output)
        }
        fmt.Printf("Using format %s\n", *format)
        for _, file := range *include {
                fmt.Printf("Including %s\n", file)
        }
}
```

```
./hello --output -f myformat Bob -I myfile -Imyotherfile --include=yetanotherfile
```