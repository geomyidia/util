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

func ParseOptions(opts []Option) *Options {
	options := &Options{
		PathSeparator: defaultPathSeparator,
		Paths:         defaultPaths,
	}
	for _, opt := range opts {
		opt(options)
	}
	return options
}
