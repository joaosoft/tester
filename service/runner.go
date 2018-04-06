package gotest

type IRunner interface {
	Run() error
}

type Runner struct {
	testFiles map[string]*TestFile
	runners   []IRunner
}

func NewRunner(testFiles map[string]*TestFile) *Runner {
	return &Runner{
		testFiles: testFiles,
		runners:   make([]IRunner, 0),
	}
}

func (runner *Runner) Run() error {
	for testFileName, testFile := range runner.testFiles {
		log.Infof("running test file %s", testFileName)

		scenarioRunner := NewScenarioRunner(&testFile.Scenario)

		if err := scenarioRunner.Setup(); err != nil {
			return err
		}

		if err := testFile.Tests.Run(scenarioRunner); err != nil {
			return err
		}

		if err := scenarioRunner.Teardown(); err != nil {
			return err
		}
	}
	return nil
}
