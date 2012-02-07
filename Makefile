include $(GOROOT)/src/Make.inc

TARG=pokey
GOFILES=\
	pokey.go\
	flags.go

include $(GOROOT)/src/Make.cmd
