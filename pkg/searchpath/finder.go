package searchpath

import (
	"os"
	"path/filepath"
)

func FindFile(filename string, opts *Options) (string, error) {
	return findFile(filename, opts.Paths)
}

// Support functions

func findFile(filename string, paths []string) (string, error) {
	var fullpath string
	for _, path := range paths {
		fullpath = filepath.Join(path, filename)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, nil
		}
	}
	return "", ErrNotFound
}
