package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func buildArchiveCommand(file string) []string {
	return []string{"atool", "-l", file}
}

func buildImageCommand(file, width, height, x, y string) []string {
	switch IMAGE_PROTOCOL {
	case KITTY_KITTEN:
		return []string{
			"kitten",
			"icat",
			"--stdin", "no",
			"--transfer-mode", "memory",
			"--place", fmt.Sprintf("%sx%s@%sx%s", width, height, x, y),
			file,
		}
	case KITTY_CHAFA, SIXEL:
		format := "kitty"
		if IMAGE_PROTOCOL == SIXEL {
			format = "sixels"
		}
		return []string{
			"chafa", "-s", width + "x" + height, "--format", format,
			"--bg", "black", "--polite", "on", file,
		}
	}
	return nil
}

func buildMarkdownCommand(file string, width string) []string {
	return []string{"glow", "-w", width, file}
}

func buildTextCommand(file string, width string) []string {
	return []string{
		"bat", "--color=always", "--style=plain", "--paging=never",
		"--wrap=character", "--terminal-width", width, "--", file,
	}
}

func buildPDFCommand(file, width, height, x, y string) []string {
	clearCacheIfTooBig(getCacheDirectory(), int64(500)*1024*1024)

	cacheFile, err := makeCachePath(file)
	if err != nil {
		panic(err)
	}

	path := cacheFile[:len(cacheFile)-4]
	args := []string{
		"pdftoppm", "-f", "1", "-l", "1", "-scale-to-x", "1920",
		"-scale-to-y", "-1", "-singlefile", "-jpeg", file, path,
	}

	if !isCacheValid(file, cacheFile) {
		generateThumbnail(cacheFile, args)
	}

	return buildImageCommand(cacheFile, width, height, x, y)
}

func buildVideoCommand(file, width, height, x, y string) []string {
	clearCacheIfTooBig(getCacheDirectory(), int64(500)*1024*1024)

	cacheFile, err := makeCachePath(file)
	if err != nil {
		panic(err)
	}

	args := []string{"ffmpegthumbnailer", "-i", file, "-o", cacheFile, "-s", "0", "-t", "50%"}

	if !isCacheValid(file, cacheFile) {
		generateThumbnail(cacheFile, args)
	}

	return buildImageCommand(cacheFile, width, height, x, y)
}

func buildLibreOfficeCommand(file, width, height, x, y string) []string {
	cacheDirectory := getCacheDirectory()
	clearCacheIfTooBig(cacheDirectory, int64(500)*1024*1024)

	cacheFile, err := makeCachePath(file)
	if err != nil {
		panic(err)
	}

	if !isCacheValid(file, cacheFile) {
		args := []string{"libreoffice", "--headless", "--convert-to", "jpg", "--outdir", cacheDirectory, file}

		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = io.Discard
		cmd.Stderr = os.Stderr
		cmd.Run()

		path := filepath.Base(file)
		path = path[:len(path)-4] + "jpg" // docx is 4 chars
		path = filepath.Join(cacheDirectory, path)
		os.Rename(path, cacheFile)
	}

	return buildImageCommand(cacheFile, width, height, x, y)
}
