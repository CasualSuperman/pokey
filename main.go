package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	uid := os.Geteuid()
	var prog *exec.Cmd

	if uid == 1 {
		prog = runAsRoot()
	} else {
		prog = runAsUser()
	}

	stdout, _ := prog.StdoutPipe()
	stderr, _ := prog.StderrPipe()

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	err := prog.Run()
	if err != nil {
		fmt.Println(err.Error(), err)
	}
}

func runAsRoot() *exec.Cmd {
	return exec.Command("pacman", os.Args[1:]...)
}

func runAsUser() *exec.Cmd {
	return exec.Command("sudo", append([]string{"--", "pacman"}, os.Args[1:]...)...)
}
