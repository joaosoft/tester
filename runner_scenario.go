package tester

import (
	"github.com/joaosoft/logger"
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
	logger    logger.ILogger
	runner    *Runner
}

// NewScenarioRunner ...
func (runner *Runner) NewScenarioRunner(scenario *Scenario) (*ScenarioRunner, error) {
	if scenarios, err := runner.load(scenario); err != nil {
		return nil, err
	} else {
		return &ScenarioRunner{
			scenarios: scenarios,
			logger:    runner.logger,
			runner:    runner,
		}, nil
	}
}

// load recursive load scenario files inside every scenario
func (runner *Runner) load(scenario *Scenario) ([]*Scenario, error) {
	runner.logger.Info("loading scenarios...")
	array := make([]*Scenario, 0)

	for _, file := range scenario.Files {
		runner.logger.Infof("loading scenario file %s", file)
		nextScenario := &Scenario{}
		if _, err := ReadFile(file, nextScenario); err != nil {
			return nil, err
		}

		runner.logger.Infof("getting next scenario...")
		if nextArray, err := runner.load(nextScenario); err != nil {
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

	runner.logger.Info("setup scenario...")
	for _, scenario := range runner.scenarios {
		services = append(services, scenario.Setup...)
	}

	runner.setup = setup.NewSetup(setup.WithRunInBackground(true), setup.WithLogger(runner.logger), setup.WithServices(services))
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
