package tester

func (tests *Tests) Run(scenarioRunner *ScenarioRunner) error {
	tests.logger.Info("run http tests...")
	if err := scenarioRunner.runner.NewWebRunner(scenarioRunner, tests.HttpTest).Run(); err != nil {
		return err
	}

	tests.logger.Info("run sql tests...")
	if err := scenarioRunner.runner.NewSQLRunner(scenarioRunner, tests.SqlTest).Run(); err != nil {
		return err
	}

	tests.logger.Info("run redis tests...")
	if err := scenarioRunner.runner.NewRedisRunner(scenarioRunner, tests.RedisTest).Run(); err != nil {
		return err
	}

	return nil
}
