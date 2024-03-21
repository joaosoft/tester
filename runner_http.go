package tester

import (
	"fmt"
	"github.com/joaosoft/web"

	"github.com/joaosoft/logger"

	"encoding/json"

	manager "github.com/joaosoft/manager"
)

type HttpRunner struct {
	tests  []HttpTest
	pm     *manager.Manager
	logger logger.ILogger
}

func (runner *Runner) NewWebRunner(scenarioRunner *ScenarioRunner, tests []HttpTest) *HttpRunner {
	return &HttpRunner{
		tests:  tests,
		logger: runner.logger,
		pm:     runner.pm,
	}
}

func (runner *HttpRunner) Run() error {
	gateway, err := runner.pm.NewSimpleGateway()
	if err != nil {
		return fmt.Errorf("error running http tester [ error: %s ]", err)
	}

	for _, test := range runner.tests {
		runner.logger.Infof("running http tester with [ name: %s, description: %s ]", test.Name, test.Description)
		var headers HttpHeaders
		if test.Headers != nil {
			headers = *test.Headers
		}

		body, err := json.Marshal(test.Body)
		if err != nil {
			return fmt.Errorf("error executing http tester [ error: %s ]", err)
		}

		status, response, err := gateway.Request(test.Method, test.Host, test.Route, string(web.ContentTypeApplicationJSON), headers, body)

		if err != nil {
			return fmt.Errorf("error executing http tester [ error: %s ]", err)
		}

		if status != test.Expected.Status {
			return fmt.Errorf("return status should be %d instead of %d", test.Expected.Status, status)
		}

		matcher := NewMatcher(test.Expected.Body)
		if err := matcher.Match(response); err != nil {
			return err
		}
	}

	return nil
}
