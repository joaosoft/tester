package main

import (
	"os"

	"go-test/service"

	"time"

	logger "github.com/joaosoft/go-log/service"
	writer "github.com/joaosoft/go-writer/service"
)

func main() {
	var log = logger.NewLog(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFormatHandler(writer.JsonFormatHandler),
		logger.WithWriter(os.Stdout)).WithPrefixes(map[string]interface{}{
		"level":   logger.LEVEL,
		"time":    logger.TIME,
		"service": "go-test"})

	start := time.Now()
	test := gotest.NewGoTest(gotest.WithLogger(log), gotest.WithPath("./examples"))

	test.Run()

	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())
}
