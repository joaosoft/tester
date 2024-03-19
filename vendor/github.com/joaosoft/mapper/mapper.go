package mapper

import (
	gomanager "github.com/joaosoft/go-manager"
	"github.com/joaosoft/logger"
)

// Mapper ...
type Mapper struct {
	config        *MapperConfig
	pm            *gomanager.Manager
	logger        logger.ILogger
	isLogExternal bool
}

// NewMapper ...
func NewMapper(options ...MapperOption) *Mapper {
	config, simpleConfig, err := NewConfig()

	service := &Mapper{
		pm:      gomanager.NewManager(gomanager.WithRunInBackground(false)),
		logger: logger.NewLogDefault("mapper", logger.WarnLevel),
		config: config.Mapper,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(gomanager.WithLogger(log))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Mapper != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Mapper.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	service.Reconfigure(options...)

	return service
}
