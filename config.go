package tester

import (
	"fmt"

	manager "github.com/joaosoft/manager"
)

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

// NewConfig ...
func NewConfig() (*TesterConfig, error) {
	appConfig := &AppConfig{}
	if _, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		log.Error(err.Error())

		return &TesterConfig{}, err
	}

	return appConfig.Tester, nil
}
