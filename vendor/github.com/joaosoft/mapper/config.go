package mapper

import (
	"fmt"
	"github.com/joaosoft/go-manager"
)

// AppConfig ...
type AppConfig struct {
	Mapper *MapperConfig `json:"mapper"`
}

// MapperConfig ...
type MapperConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
