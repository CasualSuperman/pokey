include $(GOROOT)/src/Make.inc

TARG=pacpill
GOFILES=\
	pacpill.go\
	flags.go

include $(GOROOT)/src/Make.cmd
