package gotest

// appConfig ...
type appConfig struct {
	GoTest goTestConfig `json:"gotest"`
}

// goTestConfig ...
type goTestConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}
