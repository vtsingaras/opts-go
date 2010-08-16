// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The opts package provides advanced GNU- and POSIX- style option parsing.
*/
package opts

import (
	"fmt"
	"os"
	"strings"
)

//
// Exported variables
//

// The name with which this program was called
var Xname = os.Args[0]

// The list of optionless arguments provided
var Args []string = make([]string, 0, len(os.Args))

// A description of the program, which may be multiline
var Description string

// A string with the usage of the program. usage: and the name of the program
// are automatically prefixed.
var Usage string = "[options]"

//
// Description of options
//

// The built-in types of options.
const (
	FLAG = iota
	HALF
	SINGLE
	MULTI
)

// The built-in types of errors.
const (
	UNKNOWN = iota // unknown option
	REQARG // a required argument was not present
	NOARG // an argument was present where none should have been
)

// Parsing is a callback used by Option implementations to report errors.
type Parsing struct{}

// Error prints the relevant error message and exits.
func (Parsing) Error(err int, opt string) {
	switch err {
		case UNKNOWN:
			fmt.Fprintf(os.Stderr, 
				"%s: %s: unknown option",
				Xname, opt)
		case REQARG:
			fmt.Fprintf(os.Stderr,
				"%s: %s: required argument",
				Xname, opt)
		case NOARG:
			fmt.Fprintf(os.Stderr,
				"%s: %s takes no argument",
				Xname, opt)
			
	}
	os.Exit(1)
}

// Option represents a conceptual option, which there may be multiple methods
// of invoking.
type Option interface {
	// Forms returns a slice with all forms of this option. Forms that 
	// begin with a single dash are interpreted as short options, multiple
	// of which may be combined in one argument (-abcdef). Forms that begin
	// with two dashes are interpreted as long options, which must remain
	// whole.
	Forms() []string
	// Description returns the description of this option.
	Description() string
	// ArgName returns a descriptive name for the argument this option
	// takes, or an empty string if this option takes none.
	ArgName() string
	// Invoke is called when this option appears in the command line.
	// If the option expects an argument (as indicated by ArgName), 
	// Invoke is guaranteed not to be called without one. Similarly, if 
	// the option does not expect an argument, Invoke is guaranteed to be
	// called only with the first parameter being the empty string.
	Invoke(string, Parsing)
}

// A partial implementation of the Option interface that reflects what most
// options share.
type genopt struct {
	shortform string
	longform string
	description string
}

func (o genopt) Forms() []string {
	forms := make([]string, 0, 2)
	if len(o.shortform) > 0 {
		forms = forms[0:1]
		forms[0] = o.shortform
	}
	if len(o.longform) > 0 {
		forms = forms[0:len(forms)+1]
		forms[len(forms)-1] = o.longform
	}
	return forms
}

func (o genopt) Description() string { return o.description }

type flag struct {
	genopt
	dest *bool
}

func (flag) ArgName() string { return "" }

type half struct {
	genopt
	dest *string
	dflt string // the value if the option is not given
	givendflt string // the default value if the option is given
}

func (o half) ArgName() string { return o.givendflt }

type single struct {
	genopt
	dest *string
	dflt string
}

func (o single) ArgName() string { return o.dflt }

type multi struct {
	genopt
	valuedesc string
	dest *[]string
}

func (o multi) ArgName() string { return o.valuedesc }

// Stores an option of any kind
type option struct {
	dflt        string
	strdest     *string
	strslicedest *[]string
}

// The registered options
var options map[string]Option = map[string]Option{}

// Add adds the given option.
func Add(opt Option) {
	for _, form := range opt.Forms() {
		options[form]=opt
	}
}

// True if the option list has been terminated by '-', false otherwise.
var optsOver bool

// Parse performs parsing of the command line, making complete information 
// available to the program.
func Parse() {
	ParseArgs(os.Args)
}

// ParseArgs performs parsing of the given command line, making complete
// information available to the program.
//
// This function was created specifically to enable unit testing - the proper
// entry point for most programs is Parse.
func ParseArgs(args []string) {
	addHelp()
	p := Parsing{}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if len(arg)>0 && arg[0]=='-' && !optsOver {
			switch {
			case len(arg)==1:
				optsOver=true
			case arg[1]=='-':
				argparts := strings.Split(arg, "=", 2)
				arg, val := argparts[0], argparts[1]
				if option, ok := options[arg]; ok {
					switch {
					case val=="" && option.ArgName()=="":
					}
				} else {
					p.Error(UNKNOWN, arg)
				}
			default:
				for _, optChar := range arg[1:len(arg)] {
					opt := string(optChar)
					if option, ok := options["-"+opt]; ok {
						if option.ArgName()=="" {
							
						} else {
							
						}
					} else {
						p.Error(UNKNOWN, "-"+opt)
					}
				}
			}
		} else {
			Args = Args[0:len(Args)+1]
			Args[len(Args)-1] = arg
		}
	}
}
