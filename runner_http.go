package tester

import (
	"fmt"

	"encoding/json"

	manager "github.com/joaosoft/manager"
)

type HttpRunner struct {
	tests []HttpTest
}

func NewWebRunner(scenarioRunner *ScenarioRunner, tests []HttpTest) *HttpRunner {
	return &HttpRunner{
		tests: tests,
	}
}

func (runner *HttpRunner) Run() error {
	gateway := manager.NewSimpleGateway()

	for _, test := range runner.tests {
		log.Infof("running http tester with [ name: %s, description: %s ]", test.Name, test.Description)
		var headers HttpHeaders
		if test.Headers != nil {
			headers = *test.Headers
		}

		body, err := json.Marshal(test.Body)
		if err != nil {
			return fmt.Errorf("error executing http tester [ error: %s ]", err)
		}

		status, response, err := gateway.Request(test.Method, test.Host, test.Route, headers, body)

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