// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The opts package provides basic GNU- and POSIX- style 
// option parsing, similarly to getopt.
package opts

import (
	"container/vector"
	"fmt"
	"io"
	"os"
	"strings"
	"tabwriter"
)

// The name with which this program was called
var Xname = os.Args[0]

// The list of optionless arguments provided
var Args vector.StringVector

// A description of the program, which may be multiline
var description string

// A string with the usage of the program
var usage string = os.Args[0]+" [options]"

// Sets the program usage to the given string, prefixed with 'usage: '
func Usage(u string) {
	usage = fmt.Sprintf("usage: %s %s",Xname,u)
}

// Sets the program description to an arbitrary string.
func Description(desc string) {
	description = desc
}

type optionType int
const (
	FLAG = iota
	OPTION
	MULTI
)

// Stores an option that takes one argument ("option")
type option struct {
	optType optionType
	shortform    string
	longform     string
	description string
	dflt        string
	strdest *string
	booldest *bool
	strvecdest *vector.StringVector
}

// The registered options
var options map[string]option = map[string]option{}

// Flag creates a flag with the specified short and long forms.
func Flag(shortform string, longform string, desc string) *bool {
	dest := new(bool)
	flag := option {
		optType: FLAG,
		shortform: "-"+shortform,
		longform: "--"+longform,
		description: desc,
		dflt: "",
		booldest: dest,
	}
	// insert the items into the map
	options["-"+shortform] = flag
	options["--"+longform] = flag
	return dest
}

// Option creates an option with the specified short and long forms.
func Option(shortform string, longform string, desc string, dflt string) *string {
	dest := new(string)
	*dest = dflt
	opt := option {
		optType: OPTION,
		shortform: "-"+shortform,
		longform: "--"+longform,
		description: desc,
		dflt: dflt,
		strdest: dest,
	}
	// insert the items into the map
	options["-"+shortform] = opt
	options["--"+longform] = opt
	return dest
}

// Multi creates an option that can be called multiple times.
func Multi(shortform string, longform string, desc string) *vector.StringVector {
	dest := &vector.StringVector{}
	multi := option {
		optType: MULTI,
		shortform: "-"+shortform,
		longform: "--"+longform,
		description: desc,
		dflt: "",
		strvecdest: dest,
	}
	// insert the items into the map
	options["-"+shortform] = multi
	options["--"+longform] = multi
	return dest
}

// Shortflag creates a flag with no long form, only a short one.
func Shortflag(shortform string, desc string) *bool {
	return Flag(shortform, "", desc)
}
// Longflag ceates a flag with no short form, only a long one
func Longflag(longform string, desc string) *bool {
	return Flag("", longform, desc)
}
// Shortopt creates an option with no long form.
func Shortopt(opt string, desc string, dflt string) *string {
	return Option(opt, "", desc, dflt)
}
// Longopt creates an option with no short form.
func Longopt(opt string, desc string, dflt string) *string {
	return Option("", opt, desc, dflt)
}
// Shortmulti creates an option with no long form.
func Shortmulti(opt string, desc string) *vector.StringVector {
	return Multi(opt, "", desc)
}
// Longmulti creates an option with no short form.
func Longmulti(opt string, desc string) *vector.StringVector {
	return Multi("", opt, desc)
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

func pushValue(opt string, dest *vector.StringVector, place int) {
	if place >= len(os.Args) || isOption(os.Args[place]) {
		needArgument(opt)
		os.Exit(1)
	}
	(*dest).Push(os.Args[place])
}

func handleOption(optnum int) int {
	opt := os.Args[optnum]
	if opt[1] == '-' {
		// this is a single long argument
		// get rid of any =
		_opt := strings.Split(opt,"=",2)
		opt = _opt[0]
		// check the flags list
		if option, ok := options[opt]; ok {
			switch {
			case option.optType == FLAG:
				*option.booldest = true
			case option.optType == OPTION:
				// get the next value
				if len(_opt) > 1 {
					*option.strdest = _opt[1]
				} else {
					needArgument(opt)
				}
			case option.optType == MULTI:
				// get the next value
				if len(_opt) > 1 {
					option.strvecdest.Push(_opt[1])
				} else {
					needArgument(opt)
				}
			}
			return 0
		} else {
			// This option doesn't exist
			invalidOption(opt, optnum)
		}
	} else {
		// this is a short argument
		// for each option
		for i := 1; i < len(opt); i++ {
			o := "-" + string(opt[i])
			option, ok := options[o]
			if !ok {
				invalidOption(o, optnum)
			}
			switch {
			case option.optType == FLAG:
				*option.booldest = true
			case option.optType == OPTION && i == len(opt)-1:
				assignValue(o, option.strdest, optnum+1)
				return 1
			case option.optType == OPTION && i != len(opt)-1:
				needArgument(o)
			case option.optType == MULTI && i == len(opt)-1:
				pushValue(o, option.strvecdest, optnum+1)
				return 1
			case option.optType == MULTI && i != len(opt)-1:
				needArgument(o)
			}
		}
	}
	return 0
}

var getHelp *bool

// addHelp() the -h and --help options, if neither has been added yet. This is 
// called automatically when Parse() is called.
func addHelp() {
	_, ok := options["-h"]; _, ok2 := options["--help"]
	if !(ok || ok2) {
		getHelp = Flag("h", "help", "display help screen")
	} else {
		getHelp = new(bool)
	}
}

// Parse performs POSIX and GNU option parsing, based on previously set settings
func Parse() {
	addHelp() // If not already done, add the help option
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
	// check if help was asked for
	if *getHelp {
		Help()
		// and exit
		os.Exit(0)
	}
}

func printOption(w io.Writer,
		shortform string, 
		longform string, 
		description string, 
		dflt string,
		value bool,
		multi bool) {
	valappend := ""
	switch {
	case value && longform != "--":
		valappend = "=STRING"
	case value && longform == "--":
		valappend = " STRING"
	}
	if multi {
		valappend += " ..."
	}
	switch {
	case shortform != "-" && longform != "--":
		fmt.Fprintf(w," %s,\t%s%s\t%s",shortform,
			longform,valappend,description)
	case shortform != "-" && longform == "--":
		fmt.Fprintf(w," %s%s\t\t%s",shortform,valappend,description)
	case shortform == "-" && longform != "--":
		fmt.Fprintf(w," \t%s%s\t%s",longform,valappend,description)
	}
	// TODO FIXME print the default
	fmt.Fprintf(w,"\n")
}

// Help prints a generated help screen, from the options previously passed
func Help() {
	fmt.Printf("%s\n%s\n",usage,description)
	// a record of which options we've already printed
	done := map[string]bool{}
	// start formatting with the tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', 0)
	for str, opt := range options {
		if !done[str] {
			printOption(w,
				opt.shortform,opt.longform,
				opt.description,
				opt.dflt,opt.optType!=FLAG,
				opt.optType==MULTI)
		}
		done[opt.shortform], done[opt.longform] = true, true
	}
	// flush the tabwriter
	w.Flush()
}
