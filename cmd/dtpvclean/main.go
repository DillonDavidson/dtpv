package main

import (
	"fmt"
	"os"
	"os/exec"
)

type ImageProtocol int

const (
	SIXEL ImageProtocol = iota
	KITTY_CHAFA
	KITTY_KITTEN
)

func detectProtocol() ImageProtocol {
	if isKittyTerminal() {
		if _, err := exec.LookPath("kitten"); err == nil {
			// return KITTY_KITTEN
		}
		return KITTY_CHAFA
	}
	return SIXEL
}

func isKittyTerminal() bool {
	if os.Getenv("KITTY_WINDOW_ID") != "" {
		return true
	}

	// Only supporting ghostty and wezterm right now
	switch os.Getenv("TERM_PROGRAM") {
	case "ghostty", "WezTerm":
		return true
	}

	// if os.Getenv("TERM") == "xterm-kitty" {
	// 	return true
	// }

	return false
}

var IMAGE_PROTOCOL ImageProtocol = detectProtocol()

func buildCleanCommand() []string {
	return []string{
		"kitten",
		"icat",
		"--clear",
		"--stdin", "no",
		"--transfer-mode", "memory",
	}
}

func clearKitty() {
	tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	if err != nil {
		return
	}
	defer tty.Close()
	fmt.Fprint(tty, "\x1b_Ga=d,d=A,q=2\x1b\\")
}

func main() {
	if IMAGE_PROTOCOL == SIXEL {
		os.Exit(0)
	}

	if IMAGE_PROTOCOL == KITTY_CHAFA {
		clearKitty()
		os.Exit(0)
	}

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
