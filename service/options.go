package gotest

import (
	"strings"

	logger "github.com/joaosoft/go-log/service"
)

// GoMockOption ...
type GoMockOption func(gotest *GoMock)

// Reconfigure ...
func (gotest *GoMock) Reconfigure(options ...GoMockOption) {
	for _, option := range options {
		option(gotest)
	}
}

// WithPath ...
func WithPath(path string) GoMockOption {
	return func(gotest *GoMock) {
		if path != "" {
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			global["path"] = path
		}
	}
}

// WithRunInBackground ...
func WithRunInBackground(background bool) GoMockOption {
	return func(gotest *GoMock) {
		gotest.background = background
	}
}

// WithConfigurationFile ...
func WithConfigurationFile(file string) GoMockOption {
	return func(gotest *GoMock) {
		config := &Configurations{}
		if _, err := readFile(file, config); err != nil {
			panic(err)
		}
		gotest.Reconfigure(
			WithSQLConfiguration(&config.Connections.SQLConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNSQConfiguration(&config.Connections.NSQConfig))
	}
}

// WithRedisConfiguration ...
func WithRedisConfiguration(config *RedisConfig) GoMockOption {
	return func(gotest *GoMock) {
		global["redis"] = config
	}
}

// WithSQLConfiguration ...
func WithSQLConfiguration(config *SQLConfig) GoMockOption {
	return func(gotest *GoMock) {
		global["sql"] = config
	}
}

// WithNSQConfiguration ...
func WithNSQConfiguration(config *NSQConfig) GoMockOption {
	return func(gotest *GoMock) {
		global["nsq"] = config
	}
}

// WithLogger ...
func WithLogger(logger logger.Log) GoMockOption {
	return func(gomock *GoMock) {
		log = logger
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) GoMockOption {
	return func(gomock *GoMock) {
		log.SetLevel(level)
	}
}


// WithConfigurations ...
func WithConfigurations(config *Configurations) GoMockOption {
	return func(gotest *GoMock) {
		gotest.Reconfigure(
			WithSQLConfiguration(&config.Connections.SQLConfig),
			WithRedisConfiguration(&config.Connections.RedisConfig),
			WithNSQConfiguration(&config.Connections.NSQConfig))
	}
}
