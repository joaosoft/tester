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
			global[path_key] = path
		}
	}
}

// WithLogger ...
func WithLogger(logger logger.ILog) GoTestOption {
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
