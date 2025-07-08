// Package searchpath provides functionality to find files in specified directories.
// It allows searching for files in a list of paths, which can be specified as a string
package searchpath

import (
	"os"
	"path/filepath"

	"github.com/geomyidia/util/pkg/errors"
)

// FindFile searches for a file in the specified search paths and returns its full path.
func FindFile(filename string, opts ...Option) (string, error) {
	options := ParseOptions(opts)
	return findFile(filename, options.Paths)
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
	return "", errors.ErrFileNotFoundAny
}
