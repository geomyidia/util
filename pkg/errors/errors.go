package errors

import (
	"errors"
	"fmt"
)

var (
	// Errors
	ErrBuildInfo       = errors.New("could not load build info")
	ErrFileNotFoundAny = errors.New("given filename not found in any of the search paths")
	ErrVersionData     = errors.New("failed to get version data")
)

func Combine(err1, err2 error) error {
	return fmt.Errorf("%s (%s)", err1.Error(), err2.Error())
}
