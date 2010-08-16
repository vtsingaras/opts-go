# Copyright 2010 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include ${GOROOT}/src/Make.$(GOARCH)

TARG = opts
GOFMT = gofmt -w

GOFILES = opts.go help.go
include $(GOROOT)/src/Make.pkg

opts.6: _go_.6
	cp _go_.6 opts.6

sample: sample.go opts.6

	6g -I . sample.go
	6l -o sample sample.6

format:
	${GOFMT} ${GOFILES}
