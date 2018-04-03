package gotest

import (
	"time"

	gomock "github.com/joaosoft/go-mock/service"
)

// TestFile ...
type TestFile struct {
	Name        string   `json:"name`
	Description string   `json:"description"`
	Scenario    Scenario `json:"scenario"`
	Tests       Tests    `json:"testFiles"`
}

// Scenario ...
type Scenario struct {
	Options map[string]string  `json:"options,omitempty"`
	Files   []string           `json:"files,omitempty"`
	Mocks   []*gomock.Services `json:"mocks,omitempty"`
	Tests   []Tests            `json:"testFiles"`
}

// Tests ...
type Tests struct {
	Description string      `json:"description"`
	HttpTest    []HttpTest  `json:"http"`
	SqlTest     []SqlTest   `json:"sql"`
	RedisTest   []RedisTest `json:"redis"`
}

type BaseTest struct {
	Scenario    Scenario `json:"scenario"`
	Description string   `json:"description"`
}

// HttpTest ...
type HttpTest struct {
	BaseTest
	Host    string         `json:"host"`
	Method  string         `json:"method"`
	Route   string         `json:"route"`
	Headers *HttpHeaders   `json:"headers"`
	Cookies []*HttpCookies `json:"cookies"`
	Body    struct {
	} `json:"body"`
	Expected struct {
		Status int      `json:"status"`
		Body   HttpBody `json:"body"`
	} `json:"expected"`
}

// HttpHeaders ...
type HttpHeaders map[string][]string

// HttpCookies ...
type HttpCookies struct {
	Name    *string    `json:"name"`
	Value   *string    `json:"value"`
	Path    *string    `json:"path"`    // optional
	Domain  *string    `json:"domain"`  // optional
	Expires *time.Time `json:"expires"` // optional
}

// HttpBody ...
type HttpBody struct {
	BodyMatch
}

// BodyMatch ...
type BodyMatch struct {
	Match int         `json:"matcher"`
	Value interface{} `json:"value"`
}

// SqlTest ...
type SqlTest struct {
	BaseTest
	Configuration *SqlConfig `json:"configuration"`
	Connection    *string    `json:"connection"`
	Query         string     `json:"query"`
	Expected      string     `json"expected"`
}

// SqlConfig ...
type SqlConfig struct {
	Driver     string `json:"driver"`
	DataSource string `json:"datasource"`
}

// RedisTest ...
type RedisTest struct {
	BaseTest
	Configuration *RedisConfig `json:"configuration"`
	Connection    *string      `json:"connection"`
	Command       string       `json:"command"`
	Expected      string       `json"expected"`
}

// RedisConfig ...
type RedisConfig struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Size     int    `json:"size"`
}
