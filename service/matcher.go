package gotest

import (
	"fmt"

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

	expected_mapped, _ := mapper.Map(matcher.expected)
	fmt.Printf("EXPECTED: %+v", expected_mapped)

	obtained_mapped, _ := mapper.Map(obtained)
	fmt.Printf("OBTAINED: %+v", obtained_mapped)

	return nil
}
