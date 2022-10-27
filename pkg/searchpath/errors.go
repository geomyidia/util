package searchpath

import "errors"

var (
	// Errors
	ErrNotFound = errors.New("given filename not found in any of the search paths")
)
