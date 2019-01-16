package tester

// AppConfig ...
type AppConfig struct {
	Tester *TesterConfig `json:"tester"`
}

// TesterConfig ...
type TesterConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}
