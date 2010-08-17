// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opts

import (
	"fmt"
	"os"
	"strings"
	"tabwriter"
)

// addHelp adds the -h and --help options, if they do not already exist.
func addHelp() {

}

type helpWriter struct {
	content string
}

func (w *helpWriter) Write(data []byte) (n int, err os.Error) {
	n = len(data)
	w.content += string(data)
	return
}

func optionHelp(opt Option) (str string) {
	return
}

func helpLines() (lines []string) {
	hw := &helpWriter{}
	// start formatting with the tabwriter
	w := tabwriter.NewWriter(hw, 0, 2, 1, ' ', 0)
	lines = strings.Split(hw.content, "\n", -1)
	w.Flush()
	return
}

// Help prints a generated help screen, from the options previously passed
func Help() {
	fmt.Printf("%s\n%s\n", Usage, Description)
	// a record of which options we've already printed
	done := map[string]bool{}
	for name, opt := range options {
		if !done[name] {
			for _, form := range opt.Forms() {
				done[form] = true
			}
		}
	}
}
