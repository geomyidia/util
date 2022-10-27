package searchpath

import (
	"strings"
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
