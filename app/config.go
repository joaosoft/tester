package gotest

// appConfig ...
type appConfig struct {
	GoTest TestConfig `json:"gotest"`
}

// TestConfig ...
type TestConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}
