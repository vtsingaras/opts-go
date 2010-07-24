include ${GOROOT}/src/Make.$(GOARCH)

TARG = opts
GOFMT = gofmt -w

GOFILES = opts.go
include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} opts.go
	${GOFMT} examples/hello.go
