include $(GOROOT)/src/Make.inc

TARG=pokey
GOFILES=\
	main.go\
	flags.go

include $(GOROOT)/src/Make.cmd
