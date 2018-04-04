package gotest

import (
	"fmt"

	gomanager "github.com/joaosoft/go-manager/service"
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
	gateway := gomanager.NewSimpleGateway()
	for _, test := range runner.tests {
		log.Infof("running http test [ name: %s, description: %s ]", test.Name, test.Description)
		var headers *HttpHeaders
		if test.Headers != nil {
			headers = test.Headers
		}
		status, response, err := gateway.Request(test.Method, test.Host, test.Route, *headers, test.Body)

		if err != nil {
			return fmt.Errorf("error executing http test \"%s\"", test.Description)
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
