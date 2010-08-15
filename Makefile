# Copyright 2010 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include ${GOROOT}/src/Make.$(GOARCH)

TARG = opts
GOFMT = gofmt -w

GOFILES = opts.go
include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} opts.go
