package gotest

import (
	gomapper "github.com/joaosoft/go-mapper/service"
)

type IMatcher interface {
	Match(obtained interface{}) error
}

type Matcher struct {
	expected interface{}
}

func NewMatcher(expected interface{}) IMatcher {
	return &Matcher{
		expected: expected,
	}
}

func (matcher Matcher) Match(obtained interface{}) error {
	mapper := gomapper.NewMapper()
	return nil
}
