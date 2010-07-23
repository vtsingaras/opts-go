package opts

import (
	"fmt"
	"os"
)

// Stores an option that takes no arguments ("flag")
type FlagOpt struct {
	shortflag string
	longflag string
	description string
}

// Stores an option that takes one argument ("option")
type VarOpt struct {
	shortopt string
	longopt string
	description string
	dflt string
}

// The name with which this program was called
var Xname string

// Creates a flag with the specified short and long forms
func Flag(shortflag string, longflag string, desc string) {

}

// Creates a flag with no long form, only a short one
func Shortflag(shortflag string, desc string) {
	Flag(shortflag, "", desc)
}

// Creates a flag with no short form, only a long one
func Longflag(longflag string, desc string) {
	Flag("", longflag, desc)
}

// Creates an option with the specified short and long forms
func Option(shortopt string,longopt string,desc string,dflt string) {

}

// Creates an option with no long form
func Shortopt(opt string,desc string,dflt string) {
	Option(opt,"",desc,dflt)
}

// Creates an option with no short form
func Longopt(opt string,desc string,dflt string) {
	Option("",opt,desc,dflt)
}

// Performs POSIX and GNU option parsing, based on previously set settings
func Parse() {
	Xname=os.Args[0]
	// for each argument
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("argument: %s\n",os.Args[i]);
	}
}

