package gotest

import (
	gomock "github.com/joaosoft/go-mock/service"
)

type IScenario interface {
	Setup() error
	Teardown() error
	IsToRunOnce() bool
}

// Setup ..
func (scenario *Scenario) Setup() error {
	// scenario by file
	for _, file := range scenario.Files {
		scenarioOnFile := &Scenario{}
		if _, err := readFile(file, scenarioOnFile); err != nil {
			return err
		}
		if err := scenarioOnFile.Setup(); err != nil {
			return err
		}
	}

	// scenario by service
	for _, test := range scenario.Tests {
		if err := test.Run(); err != nil {
			return err
		}
	}

	// mocks
	runner := gomock.NewRunner(scenario.Mocks)
	runner.Setup()

	return nil
}

// Teardown ...
func (scenario *Scenario) Teardown() error {
	// files
	for _, file := range scenario.Files {
		scenarioOnFile := &Scenario{}
		if _, err := readFile(file, scenarioOnFile); err != nil {
			return err
		}
		if err := scenarioOnFile.Teardown(); err != nil {
			return err
		}
	}

	// testFiles
	for _, test := range scenario.Tests {
		if err := test.Run(); err != nil {
			return err
		}
	}

	// mocks
	runner := gomock.NewRunner(scenario.Mocks)
	runner.Setup()

	return nil
}

func (scenario *Scenario) IsToRunOnce() bool {
	if value, ok := scenario.Options["run"]; !ok || value != CONST_RUN_ONCE {
		return false
	}
	return true
}
