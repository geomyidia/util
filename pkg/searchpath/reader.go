package searchpath

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	defaultPathSeparator = ":"
)

var (
	defaultPaths = []string{"."}
	// Errors
	ErrNotFound = errors.New("given filename not found in any of the search paths")
)

func ReadFile(filename string, opts ...Option) ([]byte, error) {
	options := &Options{
		PathSeparator: defaultPathSeparator,
		Paths:         defaultPaths,
	}
	for _, opt := range opts {
		opt(options)
	}
	fullpath, err := FindFile(filename, options)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(fullpath)
}

func FindFile(filename string, opts Options) (string, error) {
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
