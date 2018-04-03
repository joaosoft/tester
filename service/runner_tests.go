package gotest

func (tests *Tests) Run(scenario IScenario) error {

	// run http
	httpRunner := NewWebRunner(scenario, tests.HttpTest)
	httpRunner.Run()

	// run sql
	sqlRunner := NewSQLRunner(scenario, tests.SqlTest)
	sqlRunner.Run()

	// run redis
	redisRunner := NewRedisRunner(scenario, tests.RedisTest)
	redisRunner.Run()

	return nil
}
