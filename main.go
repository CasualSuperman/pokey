package main

import (
	"fmt"
	"flag"
	"os/exec"
)

flag.Usage = func() {

}

func main() {
	fmt.Println("Runing pacman.")
	flag.Parse()
	exec.Command("pacman", flag.Args()...)
}
