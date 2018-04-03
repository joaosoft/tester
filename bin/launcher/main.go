package main

import (
	"os"

	"go-test/service"

	logger "github.com/joaosoft/go-log/service"
)

func main() {
	var log = logger.NewLog(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFormatHandler(logger.JsonFormatHandler),
		logger.WithWriter(os.Stdout)).WithPrefixes(map[string]interface{}{
		"level":   logger.LEVEL,
		"time":    logger.TIME,
		"service": "go-test"})

	gotest.NewGoTest(gotest.WithLogger(log))
}
