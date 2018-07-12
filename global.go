package tester

import (
	logger "github.com/joaosoft/logger"
)

var global = make(map[string]interface{})
var log = logger.NewLogDefault("tester", logger.InfoLevel)

func init() {
	global[path_key] = defaultPath
}
