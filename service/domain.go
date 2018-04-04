package gotest

import (
	"time"

	"github.com/joaosoft/go-setup/service"
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
	Options map[string]string   `json:"options,omitempty"`
	Files   []string            `json:"files,omitempty"`
	Setup   []*gosetup.Services `json:"setup,omitempty"`
	Tests   []Tests             `json:"testFiles"`
}

// Tests ...
type Tests struct {
	BaseTest
	HttpTest  []HttpTest  `json:"http"`
	SqlTest   []SqlTest   `json:"sql"`
	RedisTest []RedisTest `json:"redis"`
	NsqTest   []NsqTest   `json:"nsq"`
}

type BaseTest struct {
	Name        string `json"name"`
	Description string `json:"description"`
}

// HttpTest ...
type HttpTest struct {
	BaseTest
	Scenario Scenario       `json:"scenario"`
	Host     string         `json:"host"`
	Method   string         `json:"method"`
	Route    string         `json:"route"`
	Headers  *HttpHeaders   `json:"headers"`
	Cookies  []*HttpCookies `json:"cookies"`
	Body     struct {
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
	Scenario      Scenario   `json:"scenario"`
	Configuration *SqlConfig `json:"configuration"`
	Connection    *string    `json:"connection"`
	Expected      SqlCommand `json"expected"`
}

// SqlCommand ...
type SqlCommand struct {
	Command *string `json:"command"`
	File    *string `json:"file"`
}

// SqlConfig ...
type SqlConfig struct {
	Driver     string `json:"driver"`
	DataSource string `json:"datasource"`
}

// RedisTest ...
type RedisTest struct {
	BaseTest
	Scenario      Scenario     `json:"scenario"`
	Configuration *RedisConfig `json:"configuration"`
	Connection    *string      `json:"connection"`
	Expected      RedisCommand `json"expected"`
}

// RedisConfig ...
type RedisConfig struct {
	Protocol string `json:"protocol"`
	Address  string `json:"address"`
	Size     int    `json:"size"`
}

type RedisCommand struct {
	File      *string  `json:"file"`
	Command   *string  `json:"command"`
	Arguments []string `json:"arguments"`
}

// NsqTest ...
type NsqTest struct {
	BaseTest
	Scenario      Scenario   `json:"scenario"`
	Configuration *NsqConfig `json:"configuration"`
	Expected      NsqCommand `json"expected"`
}

// NsqConfig ...
type NsqConfig struct {
	Lookupd      string `json:"lookupd"`
	RequeueDelay int64  `json:"requeue_delay"`
	MaxInFlight  int    `json:"max_in_flight"`
	MaxAttempts  uint16 `json:"max_attempts"`
}

type NsqCommand struct {
	Topic   string  `json:"topic"`
	File    *string `json:"file"`
	Message *[]byte `json:"message"`
}
