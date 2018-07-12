package tester

func (tests *Tests) Run(scenarioRunner *ScenarioRunner) error {
	log.Info("run http tests...")
	if err := NewWebRunner(scenarioRunner, tests.HttpTest).Run(); err != nil {
		return err
	}

	log.Info("run sql tests...")
	if err := NewSQLRunner(scenarioRunner, tests.SqlTest).Run(); err != nil {
		return err
	}

	log.Info("run redis tests...")
	if err := NewRedisRunner(scenarioRunner, tests.RedisTest).Run(); err != nil {
		return err
	}

	return nil
}
