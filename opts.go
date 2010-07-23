/*
	Basic GNU and POSIX option parsing
*/

package opts

import (
	"container/vector"
	"fmt"
	"os"
)

// Stores an option that takes no arguments ("flag")
type FlagOpt struct {
	shortflag   string
	longflag    string
	description string
	destination *bool
}

// Stores an option that takes one argument ("option")
type VarOpt struct {
	shortopt    string
	longopt     string
	description string
	dflt        string
	destination *string
}

// The registered flags
var flags map[string]FlagOpt = map[string]FlagOpt{}

// The registered options
var options map[string]VarOpt = map[string]VarOpt{}

// The name with which this program was called
var Xname string

// The list of optionless Argument provided
var Args vector.StringVector

// Creates a flag with the specified short and long forms
func Flag(shortflag string, longflag string, desc string) *bool {
	dest := new(bool)
	flag := FlagOpt{shortflag,longflag,desc,dest}
	// insert the items into the map
	flags[shortflag] = flag
	flags[longflag] = flag
	return dest
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
func Option(shortopt string, longopt string, desc string, dflt string) *string {
	dest := new(string)
	opt := VarOpt{shortopt,longopt,desc,dflt,dest}
	// insert the items into the map
	options[shortopt] = opt
	options[longopt] = opt
	return dest
}

// Creates an option with no long form
func Shortopt(opt string, desc string, dflt string) {
	Option(opt, "", desc, dflt)
}

// Creates an option with no short form
func Longopt(opt string, desc string, dflt string) {
	Option("", opt, desc, dflt)
}

func handleArgument(arg string) {
	// push on the end of the vector
	Args.Push(arg)
}

func handleOption(opt string) {
	if opt[1] == '-' {
		// this is a single long argument

	} else {
		// this is a short argument
	}
}

// Performs POSIX and GNU option parsing, based on previously set settings
func Parse() {
	Xname = os.Args[0]
	optsover := false
	// for each argument
	for i := 1; i < len(os.Args); i++ {
		// check to see what type of argument
		arg := os.Args[i]
		if arg[0] == '-' && !optsover {
			if len(arg) == 1 {
				optsover = true
			} else {
				handleOption(arg)
			}
		} else {
			handleArgument(arg)
		}
	}
	fmt.Printf("placeholder\n")
}
