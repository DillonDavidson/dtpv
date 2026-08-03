package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
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
			return KITTY_KITTEN
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

func runCommand(args []string, isImage bool, x, y string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin, _ = os.Open(os.DevNull)

	if isImage && IMAGE_PROTOCOL != SIXEL {
		tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
		if err != nil {
			cmd.Stdout = os.Stdout
		} else {
			defer tty.Close()
			if IMAGE_PROTOCOL == KITTY_CHAFA {
				time.Sleep(112 * time.Millisecond)
				fmt.Fprintf(tty, "\x1b[%s;%sH", incr(y), incr(x)) // 1-indexed
			}
			cmd.Stdout = tty
		}
	} else {
		cmd.Stdout = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		println("exec failed:", err.Error())
	}
}

func incr(s string) string {
	n, _ := strconv.Atoi(s)
	return strconv.Itoa(n + 1)
}

var IMAGE_PROTOCOL ImageProtocol = detectProtocol()

func main() {
	if len(os.Args) < 2 {
		return
	}

	file := os.Args[1]
	width := "80"
	if len(os.Args) > 2 {
		width = os.Args[2]
	}
	height := "40"
	if len(os.Args) > 3 {
		height = os.Args[3]
	}
	x := "0"
	if len(os.Args) > 4 {
		x = os.Args[4]
	}
	y := "0"
	if len(os.Args) > 5 {
		y = os.Args[5]
	}

	var args []string
	var isImage bool = true

	switch filepath.Ext(file) {
	case ".pdf":
		args = buildPDFCommand(file, width, height, x, y)
	case ".md":
		args = buildMarkdownCommand(file, width)
		isImage = false
	case ".mp4":
		fallthrough
	case ".mkv":
		args = buildVideoCommand(file, width, height, x, y)
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	case ".png":
		fallthrough
	case ".webp":
		args = buildImageCommand(file, width, height, x, y)
	case ".gz":
		fallthrough
	case ".tar":
		fallthrough
	case ".zip":
		args = buildArchiveCommand(file)
		isImage = false
	case ".pptx":
		fallthrough
	case ".ppt":
		fallthrough
	case ".docx":
		fallthrough
	case ".doc":
		args = buildLibreOfficeCommand(file, width, height, x, y)
	default:
		args = buildTextCommand(file, width)
		isImage = false
	}

	if len(args) == 0 {
		os.Exit(0)
	}

	runCommand(args, isImage, x, y)

	switch IMAGE_PROTOCOL {
	case KITTY_KITTEN:
		fallthrough
	case KITTY_CHAFA:
		os.Exit(1)
	case SIXEL:
		os.Exit(0)
	}
}
