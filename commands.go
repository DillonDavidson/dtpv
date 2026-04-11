package main

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

	cache, err := makeCachePath(file)
	if err != nil {
		panic(err)
	}

	path := cache[:len(cache)-4]
	args := []string{
		"pdftoppm", "-f", "1", "-l", "1", "-scale-to-x", "1920",
		"-scale-to-y", "-1", "-singlefile", "-jpeg", file, path,
	}

	if !isCacheValid(file, cache) {
		generateThumbnail(cache, args)
	}

	return []string{"chafa", "-s", width + "x" + height, "-f", "sixels", "--bg", "black", "--polite", "on", cache}
}

func buildVideoCommand(file string, width string, height string) []string {
	clearCacheIfTooBig(getCacheDirectory(), int64(500)*1024*1024)

	cache, err := makeCachePath(file)
	if err != nil {
		panic(err)
	}

	args := []string{"ffmpegthumbnailer", "-i", file, "-o", cache, "-s", "0", "-t", "50%"}

	if !isCacheValid(file, cache) {
		generateThumbnail(cache, args)
	}

	return []string{"chafa", "-s", width + "x" + height, "-f", "sixels", "--bg", "black", "--polite", "on", cache}
}
