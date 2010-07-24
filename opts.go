// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// A description of the program, which may be multiline
var description string

// A string with the usage of the program
var usage string = os.Args[0]+" [options]"

// The list of optionless arguments provided
var Args vector.StringVector

// Sets the program usage to the given string, prefixed with 'usage: '
func Usage(u string) {
	usage = "usage: "+u
}

// Sets the program description to an arbitrary string.
func Description(desc string) {
	description = desc
}

// Creates a flag with the specified short and long forms
func Flag(shortflag string, longflag string, desc string) *bool {
	dest := new(bool)
	flag := FlagOpt{"-" + shortflag, "--" + longflag, desc, dest}
	// insert the items into the map
	flags["-"+shortflag] = flag
	flags["--"+longflag] = flag
	return dest
}

// Creates a flag with no long form, only a short one
func Shortflag(shortflag string, desc string) *bool {
	return Flag(shortflag, "", desc)
}

// Creates a flag with no short form, only a long one
func Longflag(longflag string, desc string) *bool {
	return Flag("", longflag, desc)
}

// Creates an option with the specified short and long forms
func Option(shortopt string, longopt string, desc string, dflt string) *string {
	dest := new(string)
	*dest = dflt
	opt := VarOpt{"-" + shortopt, "--" + longopt, desc, dflt, dest}
	// insert the items into the map
	options["-"+shortopt] = opt
	options["--"+longopt] = opt
	return dest
}

// Creates an option with no long form
func Shortopt(opt string, desc string, dflt string) *string {
	return Option(opt, "", desc, dflt)
}

// Creates an option with no short form
func Longopt(opt string, desc string, dflt string) *string {
	return Option("", opt, desc, dflt)
}

func invalidOption(opt string, optnum int) {
	fmt.Printf("Unknown option: %s\n", opt)
	os.Exit(1)
}

func needArgument(opt string) {
	fmt.Printf("Argument required: %s\n", opt)
	os.Exit(1)
}

var optsover bool

func isOption(opt string) bool {
	if opt[0] != '-' || optsover {
		return false
	}
	return true
}

func assignValue(opt string, dest *string, place int) {
	if place >= len(os.Args) || isOption(os.Args[place]) {
		needArgument(opt)
		os.Exit(1)
	}
	*dest = os.Args[place]
}

func handleOption(optnum int) int {
	opt := os.Args[optnum]
	if opt[1] == '-' {
		// this is a single long argument
		// check the flags list
		if flag, ok := flags[opt]; ok {
			*flag.destination = true
		} else if option, ok := options[opt]; ok {
			// get the next value
			assignValue(opt, option.destination, optnum+1)
			return 1
		} else {
			// This option doesn't exist
			invalidOption(opt, optnum)
		}
	} else {
		// this is a short argument
		// for each option
		for i := 1; i < len(opt); i++ {
			o := "-" + string(opt[i])
			flag, flagok := flags[o]
			option, optok := options[o]
			switch {
			case flagok:
				*flag.destination = true
			case optok && i == len(opt)-1:
				assignValue(o, option.destination, optnum+1)
				return 1
			case optok && i != len(opt)-1:
				needArgument(o)
			default:
				invalidOption(o, optnum)
			}
		}
	}
	return 0
}

// Parse performs POSIX and GNU option parsing, based on previously set settings
func Parse() {
	Xname = os.Args[0]
	// for each argument
	for i := 1; i < len(os.Args); i++ {
		// check to see what type of argument
		arg := os.Args[i]
		if arg[0] == '-' && !optsover {
			if len(arg) == 1 {
				optsover = true
			} else {
				i += handleOption(i)
			}
		} else {
			Args.Push(arg)
		}
	}
}

func printOption(shortform string, longform string, description string, dflt string) {
	switch {
	case shortform != "-" && longform != "--":
		fmt.Printf(" %s,\t%s\t%s",shortform,longform,description)
	case shortform != "-" && longform == "--":
		fmt.Printf(" %s\t\t%s",shortform,description)
	case shortform == "-" && longform != "--":
		fmt.Printf(" \t%s\t%s",longform,description)
	}
	// TODO FIXME print the default
	fmt.Printf("\n")
}

// Help prints a generated help screen, from the options previously passed
func Help() {
	fmt.Printf("%s\n%s\n",usage,description)
	// a record of which options we've already printed
	done := map[string]bool{}
	for str, flag := range flags {
		if !done[str] {
			printOption(flag.shortflag,
				flag.longflag,
				flag.description,
				"")
		}
		done[flag.shortflag], done[flag.longflag] = true, true
	}
	for str, opt := range options {
		if !done[str] {
			printOption(opt.shortopt,
				opt.longopt,
				opt.description,
				opt.dflt)
		}
		done[opt.shortopt], done[opt.longopt] = true, true
	}
	// TODO FIXME create actual help screen
}
