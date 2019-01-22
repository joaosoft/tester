package tester

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis"
	nsqlib "github.com/nsqio/go-nsq"
)

// createConnection ...
func (config *RedisConfig) connect() (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Database,
	}), nil
}

// createConnection ...
func (config *SqlConfig) connect() (*sql.DB, error) {
	return sql.Open(config.Driver, config.DataSource)
}

// createConnection ...
func (config *NsqConfig) connect() (*nsqlib.Producer, error) {
	nsqConfig := nsqlib.NewConfig()
	nsqConfig.MaxAttempts = config.MaxAttempts
	nsqConfig.DefaultRequeueDelay = time.Duration(config.RequeueDelay) * time.Second
	nsqConfig.MaxInFlight = config.MaxInFlight
	nsqConfig.ReadTimeout = 120 * time.Second

	return nsqlib.NewProducer(config.Lookupd, nsqConfig)
}
