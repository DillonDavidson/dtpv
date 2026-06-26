package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func runCommand(args []string, isImage bool) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr

	if isImage {
		tty, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
		if err != nil {
			cmd.Stdout = os.Stdout
		} else {
			defer tty.Close()
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
		return
	}

	runCommand(args, isImage)
	os.Exit(1)

	// cmd := exec.Command(args[0], args[1:]...)
	//
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	//
	// err := cmd.Run()
	// if err != nil {
	// 	println("exec failed:", err.Error())
	// }
}
