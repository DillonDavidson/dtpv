package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func getCacheDirectory() string {
	cachePath := os.Getenv("XDG_CACHE_HOME")
	if cachePath == "" {
		cachePath = os.Getenv("HOME") + "/.cache"
	}
	cachePath += "/dtpv"

	_, err := os.Stat(cachePath)
	if os.IsNotExist(err) {
		os.MkdirAll(cachePath, 0o755)
	}

	return cachePath
}

func makeCachePath(file string) (string, error) {
	info, err := os.Stat(file)
	if err != nil || !info.Mode().IsRegular() {
		return "", nil
	}

	canonicalPath, err := filepath.Abs(file)
	if err != nil {
		return "", nil
	}

	key := canonicalPath + strconv.FormatInt(info.Size(), 10) + strconv.FormatInt(info.ModTime().UnixNano(), 10)

	hashBytes := sha256.Sum256([]byte(key))
	hash := hex.EncodeToString(hashBytes[:])

	cachePath := filepath.Join(getCacheDirectory(), hash+".jpg")

	return cachePath, nil
}

func isCacheValid(src string, cache string) bool {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return false
	}

	cacheInfo, err := os.Stat(cache)
	if err != nil {
		return false
	}

	return !cacheInfo.ModTime().Before(srcInfo.ModTime())
}

func generateThumbnail(cache string, args []string) {
	lock := cache + ".lock"

	// If lock exists, bail
	if _, err := os.Stat(lock); err == nil {
		return
	}

	// Try to create lock file exclusively
	lockFile, err := os.OpenFile(lock, os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return
	}
	lockFile.Close()

	// Ensure lock cleanup
	defer os.Remove(lock)

	if len(args) == 0 {
		return
	}

	cmd := exec.Command(args[0], args[1:]...)

	err = cmd.Run()
	if err != nil {
		return
	}
}

func clearCacheIfTooBig(cacheDir string, maxSize int64) error {
	var total int64 = 0

	// First pass: calculate total size
	err := filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err == nil {
				total += info.Size()
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	if total <= maxSize {
		return nil
	}

	// Second pass: delete everything inside cacheDir
	entries, err := os.ReadDir(cacheDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(cacheDir, entry.Name())
		_ = os.RemoveAll(fullPath)
	}

	return nil
}
