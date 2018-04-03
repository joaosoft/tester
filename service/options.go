package gotest

import (
	"strings"

	logger "github.com/joaosoft/go-log/service"
)

// GoTestOption ...
type GoTestOption func(gotest *GoTest)

// Reconfigure ...
func (gotest *GoTest) Reconfigure(options ...GoTestOption) {
	for _, option := range options {
		option(gotest)
	}
}

// WithPath ...
func WithPath(path string) GoTestOption {
	return func(gotest *GoTest) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global["path"] = path
		}
	}
}

// WithRunInBackground ...
func WithRunInBackground(background bool) GoTestOption {
	return func(gotest *GoTest) {
		gotest.background = background
	}
}

// WithLogger ...
func WithLogger(logger logger.Log) GoTestOption {
	return func(gotest *GoTest) {
		log = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) GoTestOption {
	return func(gotest *GoTest) {
		log.SetLevel(level)
	}
}
