package main

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/tester"
	"time"
)

func main() {
	var log = logger.NewLogDefault("tester", logger.LevelInfo)

	start := time.Now()
	test := tester.NewTester(tester.WithPath("/example"))

	test.Run()

	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())
}
