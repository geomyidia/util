package searchpath

import (
	"os"
)

// ReadFile reads the contents of a file specified by filename from the search paths.
func ReadFile(filename string, opts ...Option) ([]byte, error) {
	fullpath, err := FindFile(filename, opts...)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(fullpath)
}
