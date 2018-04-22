package gotest

import (
	"strings"

	"github.com/joaosoft/go-log/service"
)

// TestOption ...
type TestOption func(test *Test)

// Reconfigure ...
func (test *Test) Reconfigure(options ...TestOption) {
	for _, option := range options {
		option(test)
	}
}

// WithPath ...
func WithPath(path string) TestOption {
	return func(test *Test) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global[path_key] = path
		}
	}
}

// WithLogger ...
func WithLogger(logger golog.ILog) TestOption {
	return func(test *Test) {
		log = logger
		test.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level golog.Level) TestOption {
	return func(test *Test) {
		log.SetLevel(level)
	}
}
