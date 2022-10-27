package searchpath

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultPathSeparator = ":"
)

var (
	defaultPaths = []string{"."}
	// Errors
	ErrNotFound = errors.New("given filename not found in any of the search paths")
)

// API

func ReadFile(filename string, opts ...Option) ([]byte, error) {
	options := &Options{
		PathSeparator: defaultPathSeparator,
		Paths:         defaultPaths,
	}
	for _, opt := range opts {
		opt(options)
	}
	fullpath, err := findFile(filename, options.Paths)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(fullpath)
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

// Options

type Option func(*Options)

type Options struct {
	PathSeparator string
	Paths         []string
	Errors        []error
}

func WithPathStr(paths string) Option {
	return func(opts *Options) {
		opts.Paths = append(strings.Split(paths, opts.PathSeparator), opts.Paths...)
	}
}

func WithPathSlice(paths []string) Option {
	return func(opts *Options) {
		opts.Paths = append(paths, opts.Paths...)
	}
}

func WithPaths(paths ...string) Option {
	return func(opts *Options) {
		opts.Paths = append(paths, opts.Paths...)
	}
}

func WithSeparator(sep string) Option {
	return func(opts *Options) {
		opts.PathSeparator = sep
	}
}
