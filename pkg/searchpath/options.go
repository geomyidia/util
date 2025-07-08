package searchpath

import (
	"strings"
)

const (
	defaultPathSeparator = ":"
)

var (
	defaultPaths = []string{"."}
)

// Option is a function that modifies the Options struct.
type Option func(*Options)

// Options holds the configuration for search paths.
type Options struct {
	PathSeparator string
	Paths         []string
	Errors        []error
}

// WithPathStr sets the search paths for the options using a colon-separated string.
func WithPathStr(paths string) Option {
	return func(opts *Options) {
		opts.Paths = append(strings.Split(paths, opts.PathSeparator), opts.Paths...)
	}
}

// WithPathSlice sets the search paths for the options using a slice of strings.
func WithPathSlice(paths []string) Option {
	return func(opts *Options) {
		opts.Paths = append(paths, opts.Paths...)
	}
}

// WithPaths sets the search paths for the options.
func WithPaths(paths ...string) Option {
	return func(opts *Options) {
		opts.Paths = append(paths, opts.Paths...)
	}
}

// WithSeparator sets the path separator for the options.
func WithSeparator(sep string) Option {
	return func(opts *Options) {
		opts.PathSeparator = sep
	}
}

// DefaultOptions returns a new Options instance with default values.
func DefaultOptions() *Options {
	return &Options{
		PathSeparator: defaultPathSeparator,
		Paths:         defaultPaths,
	}
}

// ParseOptions takes a slice of Option functions and applies them to the default options.
func ParseOptions(opts []Option) *Options {
	options := DefaultOptions()
	for _, opt := range opts {
		opt(options)
	}
	return options
}
