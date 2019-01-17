package tester

import (
	"database/sql"

	"time"

	redis "github.com/go-redis/redis"
	nsqlib "github.com/nsqio/go-nsq"
)

// createConnection ...
func (config *RedisConfig) connect() (*redis.Client, error) {
	log.Infof("connecting with address [ %s ], database [ %d ]", config.Address, config.Database)
	return redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Database,
	}), nil
}

// createConnection ...
func (config *SqlConfig) connect() (*sql.DB, error) {
	log.Infof("connecting with driver [ %s ] and data source [ %s ]", config.Driver, config.DataSource)
	return sql.Open(config.Driver, config.DataSource)
}

// createConnection ...
func (config *NsqConfig) connect() (*nsqlib.Producer, error) {
	nsqConfig := nsqlib.NewConfig()
	nsqConfig.MaxAttempts = config.MaxAttempts
	nsqConfig.DefaultRequeueDelay = time.Duration(config.RequeueDelay) * time.Second
	nsqConfig.MaxInFlight = config.MaxInFlight
	nsqConfig.ReadTimeout = 120 * time.Second

	log.Infof("connecting with max attempts [ %d ]", config.MaxAttempts)

	return nsqlib.NewProducer(config.Lookupd, nsqConfig)
}
