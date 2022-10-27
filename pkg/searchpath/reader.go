package searchpath

import (
	"os"
)

func ReadFile(filename string, opts ...Option) ([]byte, error) {
	fullpath, err := FindFile(filename, opts...)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(fullpath)
}
