package catalog

import (
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	path string
	ext  string
	os.FileInfo
}

func newFile(path string, info os.FileInfo) *file {
	name := info.Name()

	return &file{
		path:     path,
		ext:      strings.ToLower(filepath.Ext(name)),
		FileInfo: info,
	}
}

// Walk through the given directory and returns all files
// within it, not directories, and excluding the files with
// same extensions passed
func Catalog(p string, exts ...string) ([]*file, error) {
	photos := make([]*file, 0, 10)
	var allowedExts map[string]bool

	allowAllExts := len(exts) == 0

	if !allowAllExts {
		allowedExts = make(map[string]bool)

		for _, ext := range exts {
			allowedExts[ext] = true
		}
	}

	walkFn := func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		ext := filepath.Ext(p)

		if !allowAllExts && !allowedExts[ext] {
			return err
		}

		photos = append(photos, newFile(p, info))

		return err
	}

	err := filepath.Walk(p, walkFn)

	return photos, err
}
