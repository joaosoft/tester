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

		scenarioRunner.Setup()
		testFile.Tests.Run(scenarioRunner)
		scenarioRunner.Teardown()
	}
	return nil
}
