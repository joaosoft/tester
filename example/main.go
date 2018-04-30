package main

import (
	"go-test/app"

	"time"

	golog "github.com/joaosoft/go-log/app"
)

func main() {
	var log = golog.NewLogDefault("go-test", golog.InfoLevel)

	start := time.Now()
	test := gotest.NewGoTest(gotest.WithPath("/example"))

	test.Run()

	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())
}
