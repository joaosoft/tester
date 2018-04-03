package gotest

import "github.com/joaosoft/go-setup/service"

type ISystem interface {
	Setup() error
	Teardown() error
}

type ScenarioRunner struct {
	scenarios []*Scenario
	gosetup   *gosetup.GoSetup
}

func NewScenarioRunner(scenario *Scenario) *ScenarioRunner {
	scenarios := load(scenario)

	return &ScenarioRunner{scenarios: scenarios}
}

// load recursive load scenario files inside every scenario
func load(scenario *Scenario) []*Scenario {
	for _, file := range scenario.Files {
		nextScenario := &Scenario{}
		if _, err := readFile(file, nextScenario); err != nil {
			return nil
		}

		return append(load(nextScenario), scenario)
	}
	return make([]*Scenario, 0)
}

// Setup ...
func (runner *ScenarioRunner) Setup() error {
	var services []*gosetup.Services
	for _, scenario := range runner.scenarios {
		services = append(services, scenario.Setup...)
	}
	runner.gosetup = gosetup.NewGoSetup(gosetup.WithRunInBackground(true), gosetup.WithLogger(log), gosetup.WithServices(services))
	runner.gosetup.Run()

	return nil
}

// Teardown ...
func (runner *ScenarioRunner) Teardown() error {
	runner.gosetup.Stop()

	return nil
}

func (scenario *Scenario) IsToRunOnce() bool {
	if value, ok := scenario.Options["run"]; !ok || value != CONST_RUN_ONCE {
		return false
	}
	return true
}
