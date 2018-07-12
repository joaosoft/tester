package tester

import (
	"strings"

	logger "github.com/joaosoft/logger"
)

// TesterOption ...
type TesterOption func(tester *Tester)

// Reconfigure ...
func (tester *Tester) Reconfigure(options ...TesterOption) {
	for _, option := range options {
		option(tester)
	}
}

// WithPath ...
func WithPath(path string) TesterOption {
	return func(tester *Tester) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global[path_key] = path
		}
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) TesterOption {
	return func(tester *Tester) {
		log = logger
		tester.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) TesterOption {
	return func(tester *Tester) {
		log.SetLevel(level)
	}
}
