package tester

import (
	"fmt"
	"github.com/joaosoft/logger"
	nsqlib "github.com/nsqio/go-nsq"
)

type NSQRunner struct {
	tests         []NsqTest
	configuration *NsqConfig
	logger logger.ILogger
}

func (runner *Runner) NewNSQRunner(scenarioRunner *ScenarioRunner, services []NsqTest) *NSQRunner {
	return &NSQRunner{
		tests: services,
		logger: runner.logger,
	}
}

func (runner *NSQRunner) Run() error {
	for _, test := range runner.tests {
		runner.logger.Infof("running sql tester [ %s ] with description [ %s] ", test.Name, test.Description)

		var conn *nsqlib.Producer
		var err error
		if conn, err = test.Configuration.connect(); err != nil {
			return fmt.Errorf("failed to create nsq connection")
		}

		if test.Expected.Message != nil {
			if err := runner.runCommand(conn, test.Expected.Topic, *test.Expected.Message); err != nil {
				return err
			}

			if err := runner.runFile(conn, test.Expected.Topic, *test.Expected.File); err != nil {
				return err
			}
		}
	}
	return nil
}

func (runner *NSQRunner) runCommand(conn *nsqlib.Producer, topic string, message []byte) error {
	runner.logger.Infof("executing nsq [topic: %s, message: %s]", topic, message)
	if err := conn.Publish(topic, message); err != nil {
		return err
	}

	return nil
}

func (runner *NSQRunner) runFile(conn *nsqlib.Producer, topic string, file string) error {

	runner.logger.Infof("executing nsq commands by file [ %s ]", file)
	message, err := ReadFile(file, nil)
	if err != nil {
		return err
	}

	return runner.runCommand(conn, topic, message)
}
