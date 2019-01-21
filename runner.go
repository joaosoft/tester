package tester

import (
	"github.com/joaosoft/logger"
	"time"
)

type IRunner interface {
	Run() error
}

type Runner struct {
	testFiles map[string]*TestFile
	runners   []IRunner
	logger logger.ILogger
}

func (setup *Tester) NewRunner(testFiles map[string]*TestFile) *Runner {
	return &Runner{
		testFiles: testFiles,
		runners:   make([]IRunner, 0),
		logger: setup.logger,
	}
}

func (runner *Runner) Run() error {
	for testFileName, testFile := range runner.testFiles {
		runner.logger.Infof("running tester file %s", testFileName)

		scenarioRunner, err := runner.NewScenarioRunner(&testFile.Scenario)
		if err != nil {
			return err
		}

		if err := scenarioRunner.Setup(); err != nil {
			return err
		}
		<-time.After(time.Minute * 1)

		if err := testFile.Tests.Run(scenarioRunner); err != nil {
			return err
		}

		if err := scenarioRunner.Teardown(); err != nil {
			return err
		}
	}
	return nil
}
