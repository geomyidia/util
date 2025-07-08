// Package sig provides utilities for handling OS signals in Go applications.
package sig

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Handle sets up a signal handler that calls the provided handler function
func Handle(handler func(int, os.Signal), signals ...os.Signal) {
	signalHandler := make(chan os.Signal, 1)
	signal.Notify(signalHandler, signals...)
	s := <-signalHandler
	handler(os.Getpid(), s)
}

// WithContext returns a context that will be canceled on the specified signals.
func WithContext(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, signals...)
}

// DefaultSigsWithContext returns a context that will be canceled on SIGINT or SIGTERM signals.
func DefaultSigsWithContext() (context.Context, context.CancelFunc) {
	return WithContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
}
