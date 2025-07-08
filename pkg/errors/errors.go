// Package errors provides custom error types and utility functions for error handling in the application.
package errors

import (
	"errors"
	"fmt"
)

// Errors
var (
	ErrBuildInfo       = errors.New("could not load build info")
	ErrFileNotFoundAny = errors.New("given filename not found in any of the search paths")
	ErrVersionData     = errors.New("failed to get version data")
)

// Combine combines two errors into a single error message.
func Combine(err1, err2 error) error {
	return fmt.Errorf("%s (%s)", err1.Error(), err2.Error())
}
