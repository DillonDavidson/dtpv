package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

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

	var args []string

	switch filepath.Ext(file) {
	case ".pdf":
		args = buildPDFCommand(file, width, height)
	case ".md":
		args = buildMarkdownCommand(file, width)
	case ".mp4":
		fallthrough
	case ".mkv":
		args = buildVideoCommand(file, width, height)
	case ".jpeg":
		fallthrough
	case ".jpg":
		fallthrough
	case ".png":
		fallthrough
	case ".webp":
		args = buildImageCommand(file, width, height)
	case ".gz":
		fallthrough
	case ".tar":
		fallthrough
	case ".zip":
		args = buildArchiveCommand(file)
	case ".pptx":
		fallthrough
	case ".ppt":
		fallthrough
	case ".docx":
		fallthrough
	case ".doc":
		args = buildLibreOfficeCommand(file, width, height)
	default:
		args = buildTextCommand(file, width)
	}

	if len(args) == 0 {
		return
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		println("exec failed:", err.Error())
	}
}
