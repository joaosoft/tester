package tester

import (
	setup "github.com/joaosoft/setup"
)

// ISystem ...
type ISystem interface {
	Setup() error
	Teardown() error
}

// ScenarioRunner ...
type ScenarioRunner struct {
	scenarios []*Scenario
	setup     *setup.Setup
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
		if _, err := ReadFile(file, nextScenario); err != nil {
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
	var services []*setup.Services

	log.Info("setup scenario...")
	for _, scenario := range runner.scenarios {
		services = append(services, scenario.Setup...)
	}

	runner.setup = setup.NewSetup(setup.WithRunInBackground(true), setup.WithLogger(log), setup.WithServices(services))
	if err := runner.setup.Run(); err != nil {
		return err
	}

	return nil
}

// Teardown ...
func (runner *ScenarioRunner) Teardown() error {
	runner.setup.Stop()

	return nil
}

// IsToRunOnce ...
func (scenario *Scenario) IsToRunOnce() bool {
	if value, ok := scenario.Options["run"]; !ok || value != config_run_once {
		return false
	}
	return true
}
