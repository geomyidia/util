package searchpath

import (
	"os"
)

const (
	defaultPathSeparator = ":"
)

var (
	defaultPaths = []string{"."}
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
