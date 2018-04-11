package gotest

import (
	"github.com/joaosoft/go-setup/service"
)

// ISystem ...
type ISystem interface {
	Setup() error
	Teardown() error
}

// ScenarioRunner ...
type ScenarioRunner struct {
	scenarios []*Scenario
	gosetup   *gosetup.GoSetup
}

// NewScenarioRunner ...
func NewScenarioRunner(scenario *Scenario) (*ScenarioRunner, error) {
	if scenarios, err := load(scenario); err != nil {
		return nil, err
	} else {
		return &ScenarioRunner{scenarios: scenarios}, nil
	}
}

// load recursive load scenario files inside every scenario
func load(scenario *Scenario) ([]*Scenario, error) {
	log.Info("loading scenarios...")
	array := make([]*Scenario, 0)

	for _, file := range scenario.Files {
		log.Infof("loading scenario file %s", file)
		nextScenario := &Scenario{}
		if _, err := readFile(file, nextScenario); err != nil {
			return nil, err
		}

		log.Infof("getting next scenario...")
		if nextArray, err := load(nextScenario); err != nil {
			return nil, err
		} else {
			array = append(array, nextArray...)
		}
	}

	return append(array, scenario), nil
}

// Setup ...
func (runner *ScenarioRunner) Setup() error {
	var services []*gosetup.Services

	log.Info("setup scenario...")
	for _, scenario := range runner.scenarios {
		services = append(services, scenario.Setup...)
	}

	runner.gosetup = gosetup.NewGoSetup(gosetup.WithRunInBackground(true), gosetup.WithLogger(log), gosetup.WithServices(services))
	if err := runner.gosetup.Run(); err != nil {
		return err
	}

	return nil
}

// Teardown ...
func (runner *ScenarioRunner) Teardown() error {
	runner.gosetup.Stop()

	return nil
}

// IsToRunOnce ...
func (scenario *Scenario) IsToRunOnce() bool {
	if value, ok := scenario.Options["run"]; !ok || value != CONST_RUN_ONCE {
		return false
	}
	return true
}
