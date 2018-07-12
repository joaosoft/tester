package tester

import (
	"fmt"

	mapper "github.com/joaosoft/mapper"
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
	mapper := mapper.NewMapper()

	expected_mapped, _ := mapper.Map(matcher.expected)
	fmt.Printf("EXPECTED: %+v", expected_mapped)

	obtained_mapped, _ := mapper.Map(obtained)
	fmt.Printf("OBTAINED: %+v", obtained_mapped)

	return nil
}
