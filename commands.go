package main

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func buildArchiveCommand(file string) []string {
	return []string{"atool", "-l", file}
}

func buildImageCommand(file string, width string, height string) []string {
	return []string{
		"chafa", "-s", width + "x" + height, "-f", "sixels",
		"--bg", "black", "--polite", "on", file,
	}
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

func buildPDFCommand(file string, width string, height string) []string {
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

	return buildImageCommand(cacheFile, width, height)
}

func buildVideoCommand(file string, width string, height string) []string {
	clearCacheIfTooBig(getCacheDirectory(), int64(500)*1024*1024)

	cacheFile, err := makeCachePath(file)
	if err != nil {
		panic(err)
	}

	args := []string{"ffmpegthumbnailer", "-i", file, "-o", cacheFile, "-s", "0", "-t", "50%"}

	if !isCacheValid(file, cacheFile) {
		generateThumbnail(cacheFile, args)
	}

	return buildImageCommand(cacheFile, width, height)
}

func buildLibreOfficeCommand(file string, width string, height string) []string {
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

	return buildImageCommand(cacheFile, width, height)
}
