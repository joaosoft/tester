package gotest

import (
	"os"

	logger "github.com/joaosoft/go-log/service"
)

var global = make(map[string]interface{})
var log = logger.NewLog(
	logger.WithLevel(logger.InfoLevel),
	logger.WithFormatHandler(logger.JsonFormatHandler),
	logger.WithWriter(os.Stdout)).WithPrefixes(map[string]interface{}{
	"level":   logger.LEVEL,
	"time":    logger.TIME,
	"service": "go-test"})

func getDefaultNSQConfig() *NSQConfig {
	if value, exists := global["nsq"]; exists {
		return value.(*NSQConfig)
	}
	return nil
}

func getDefaultSQLConfig() *SQLConfig {
	if value, exists := global["sql"]; exists {
		return value.(*SQLConfig)
	}
	return nil
}

func getDefaultRedisConfig() *RedisConfig {
	if value, exists := global["redis"]; exists {
		return value.(*RedisConfig)
	}
	return nil
}
