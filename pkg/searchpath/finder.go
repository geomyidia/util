package searchpath

import (
	"os"
	"path/filepath"

	"github.com/geomyidia/util/pkg/errors"
)

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
