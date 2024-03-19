package setup

import (
	"fmt"
	"github.com/joaosoft/logger"

	manager "github.com/joaosoft/manager"
	"github.com/nsqio/go-nsq"
)

type NsqRunner struct {
	services      []NsqService
	configuration *manager.NSQConfig
	logger logger.ILogger
}

func (setup *Runner) NewNsqRunner(services []NsqService, config *manager.NSQConfig) *NsqRunner {
	return &NsqRunner{
		services:      services,
		configuration: config,
		logger:setup.logger,
	}
}

func (runner *NsqRunner) Setup() error {
	for _, service := range runner.services {
		runner.logger.Infof("creating service [ %s ] with description [ %s] ", service.Name, service.Description)

		var conn *nsq.Producer
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.Connect(); err != nil {
				return fmt.Errorf("failed to create nsq connection")
			}
		}

		if service.Run.Setup != nil {
			for _, setup := range service.Run.Setup {
				if err := runner.runCommands(conn, &setup); err != nil {
					return err
				}

				if err := runner.runCommandsFromFile(conn, &setup); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (runner *NsqRunner) Teardown() error {
	for _, service := range runner.services {
		runner.logger.Infof("teardown service [ %s ]", service.Name)

		var conn *nsq.Producer
		if configuration, err := runner.loadConfiguration(service); err != nil {
			return err
		} else {
			if conn, err = configuration.Connect(); err != nil {
				return fmt.Errorf("failed to create nsq connection")
			}
		}

		if service.Run.Teardown != nil {
			for _, teardown := range service.Run.Teardown {
				if err := runner.runCommands(conn, &teardown); err != nil {
					return err
				}

				if err := runner.runCommandsFromFile(conn, &teardown); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (runner *NsqRunner) loadConfiguration(test NsqService) (*manager.NSQConfig, error) {
	if test.Configuration != nil {
		return test.Configuration, nil
	} else if runner.configuration != nil {
		return runner.configuration, nil
	} else {
		return nil, fmt.Errorf("invalid nsq configuration")
	}
}

func (runner *NsqRunner) runCommands(conn *nsq.Producer, run *NsqRun) error {
	if run.Message != nil && string(run.Message) != "" {
		runner.logger.Infof("executing nsq [ %s ] message: %s", run.Description, string(run.Message))
		if err := conn.Publish(run.Topic, run.Message); err != nil {
			return err
		}
	}

	return nil
}

func (runner *NsqRunner) runCommandsFromFile(conn *nsq.Producer, run *NsqRun) error {

	if run.File != "" {
		runner.logger.Infof("executing nsq commands by file [ %s ]", run.File)
		command, err := ReadFile(run.File, nil)
		if err != nil {
			return err
		}

		if err := conn.Publish(run.Topic, command); err != nil {
			return err
		}
	}

	return nil
}
