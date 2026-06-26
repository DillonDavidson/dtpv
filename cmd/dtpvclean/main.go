package main

import (
	"os"
	"os/exec"
)

func buildCleanCommand() []string {
	return []string{
		"kitten",
		"icat",
		"--clear",
		"--stdin", "no",
		"--transfer-mode", "memory",
	}
}

func main() {
	args := buildCleanCommand()
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr

	tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	if err != nil {
		cmd.Stdout = os.Stdout
	} else {
		defer tty.Close()
		cmd.Stdout = tty
	}

	cmd.Stdin, _ = os.Open(os.DevNull)

	err = cmd.Run()
	if err != nil {
		println("exec failed:", err.Error())
	}
}
