package main

import (
	"github.com/joaosoft/tester"
	"time"

	logger "github.com/joaosoft/logger"
)

func main() {
	var log = logger.NewLogDefault("tester", logger.InfoLevel)

	start := time.Now()
	test := tester.NewTester(tester.WithPath("/example"))

	test.Run()

	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())
}
