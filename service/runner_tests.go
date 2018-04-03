package gotest

func (tests *Tests) Run(scenarioRunner *ScenarioRunner) error {
	NewWebRunner(scenarioRunner, tests.HttpTest).Run()
	NewSQLRunner(scenarioRunner, tests.SqlTest).Run()
	NewRedisRunner(scenarioRunner, tests.RedisTest).Run()

	return nil
}
