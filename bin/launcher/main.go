package main

import (
	gotest "go-test/service"
)

func main() {
	test := gotest.NewGoMock(
		gotest.WithPath("./examples"),
		gotest.WithRunInBackground(true))

	// web
	//test.RunSingle("001_webservices.json")
	//
	//// sql
	//configSQL := &gotest.SQLConfig{
	//	Driver:     "postgres",
	//	DataSource: "postgres://user:password@localhost:7001?sslmode=disable",
	//}
	//test.Reconfigure(gotest.WithSQLConfiguration(configSQL))
	//test.RunSingle("002_sql.json")
	//
	//// nsq
	//configNSQ := &gotest.NSQConfig{
	//	Lookupd:      "localhost:4150",
	//	RequeueDelay: 30,
	//	MaxInFlight:  5,
	//	MaxAttempts:  5,
	//}
	//test.Reconfigure(gotest.WithNSQConfiguration(configNSQ))
	//test.RunSingle("003_nsq.json")
	//
	//// redis
	//configRedis := &gotest.RedisConfig{
	//	Protocol: "tcp",
	//	Address:  "localhost:6379",
	//	Size:     10,
	//}
	//test.Reconfigure(gotest.WithRedisConfiguration(configRedis))
	//test.RunSingle("004_redis.json")

	// all
	test.Reconfigure(
		gotest.WithConfigurationFile("data/config.json"))

	test.Run()
	test.Wait()
	test.Stop()
}
