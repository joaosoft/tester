package gotest

import (
	"encoding/json"
	"go-mock/services"
)

type TestData struct {
	Description string   `json:"description"`
	Scenario    Scenario `json:"scenario"`
	Tests       []struct {
		HttpTest  HttpTest  `json:"http"`
		SqlTest   SqlTest   `json:"sql"`
		RedisTest RedisTest `json:"redis"`
	} `json:"tests"`
}

type Scenario []struct {
	Files []string `json:"files,omitempty"`
	Mocks []struct {
		Files []string `json:"files"`
		gomock.Services
	} `json:"mocks,omitempty"`
	Http struct {
		Description string `json:"description"`
		Host        string `json:"host"`
		Method      string `json:"method"`
		Route       string `json:"route"`
		Cookies     struct {
		} `json:"cookies"`
		Headers struct {
		} `json:"headers"`
		File string `json:"file"`
	} `json:"http,omitempty"`
}

type Http struct {
	Description string `json:"description"`
	Host        string `json:"host"`
	Method      string `json:"method"`
	Route       string `json:"route"`
	Cookies     struct {
	} `json:"cookies"`
	Headers struct {
	} `json:"headers"`
	Body json.RawMessage `json:"body"`
	File *string         `json:"file"`
}

type HttpTest struct {
	Description string `json:"description"`
	Host        string `json:"host"`
	Method      string `json:"method"`
	Route       string `json:"route"`
	Cookies     struct {
	} `json:"cookies"`
	Headers struct {
	} `json:"headers"`
	Body struct {
	} `json:"body"`
	Expected struct {
		Status int `json:"status"`
		Body   struct {
			Text string `json:"text"`
		} `json:"body"`
	} `json:"expected"`
}

type SqlTest struct {
	Description string `json:"description"`
	Host        string `json:"host"`
	Method      string `json:"method"`
	Route       string `json:"route"`
	Cookies     struct {
	} `json:"cookies"`
	Headers struct {
	} `json:"headers"`
	Body struct {
	} `json:"body"`
	Expected struct {
		Status int `json:"status"`
		Body   struct {
			Text string `json:"text"`
		} `json:"body"`
	} `json:"expected"`
}

type RedisTest struct {
	Description string `json:"description"`
	Host        string `json:"host"`
	Method      string `json:"method"`
	Route       string `json:"route"`
	Cookies     struct {
	} `json:"cookies"`
	Headers struct {
	} `json:"headers"`
	Body struct {
	} `json:"body"`
	Expected struct {
		Status int `json:"status"`
		Body   struct {
			Text string `json:"text"`
		} `json:"body"`
	} `json:"expected"`
}
