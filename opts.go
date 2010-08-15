// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	The opts package provides advanced GNU- and POSIX- style
	option parsing.
*/
package opts

import (
	"fmt"
	"os"
)

//
// Exported variables
//

// The name with which this program was called
var Xname = os.Args[0]

// The list of optionless arguments provided
var Args []string

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
	// Forms returns a slice with all forms of this option. Forms that do 
	// not begin with a dash are interpreted as short options, multiple of
	// which may be combined in one argument (-abcdef). Forms that begin
	// with a dash are interpreted as long options, which must remain
	// whole.
	Forms() []string
	// Description returns the description of this option.
	Description() string
	// ArgName returns a descriptive name for the argument this option
	// takes, or nil if this option takes none.
	ArgName() string
	// Invoke is called when this option appears in the command line.
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

func (o genopt) Description() string {
	return o.description
}

type flag struct {
	genopt
	dest *bool
}

type half struct {
	genopt
	dest *string
	dflt string // the value if the option is not given
	givendflt string // the default value if the option is given
}

type single struct {
	genopt
	dest *string
	dflt string
}

type multi struct {
	genopt
	dest *[]string
}

// Stores an option of any kind
type option struct {
	dflt        string
	strdest     *string
	strslicedest *[]string
}

// The registered options
var options map[string]Option = map[string]Option{}


