package main

import (
	"go-test/service"

	"time"

	logger "github.com/joaosoft/go-log/service"
)

func main() {
	var log = logger.NewLogDefault("go-test", logger.InfoLevel)

	start := time.Now()
	test := gotest.NewGoTest(gotest.WithLogger(log), gotest.WithPath("./examples"))

	test.Run()

	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())
}
