package setup

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// Setup ...
type Setup struct {
	services            []*Services
	runner              IRunner
	isToRunInBackground bool
	config              *SetupConfig
	pm                  *manager.Manager
	logger              logger.ILogger
	isLogExternal       bool
}

// NewSetup ...make
func NewSetup(options ...SetupOption) *Setup {
	config, simpleConfig, err := NewConfig()

	service := &Setup{
		pm:                  manager.NewManager(manager.WithRunInBackground(false)),
		logger:              logger.NewLogDefault("setup", logger.WarnLevel),
		isToRunInBackground: background,
		services:            make([]*Services, 0),
		config:              config.Setup,
	}

	service.logger.Info("starting Setup Service")

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Setup != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Setup.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))

	}

	service.Reconfigure(options...)

	return service
}

// Run ...
func (setup *Setup) Run() error {
	files, err := filepath.Glob(global[path_key].(string) + "*.json")
	if err != nil {
		return err
	}
	if err := setup.execute(files); err != nil {
		setup.logger.Error(err)
		return err
	}

	return nil
}

// RunSingle ...
func (setup *Setup) RunSingle(file string) error {
	if err := setup.execute([]string{file}); err != nil {
		setup.logger.Error(err)
		return err
	}

	return nil
}

// Stop ...
func (setup *Setup) Stop() error {
	if err := setup.runner.Teardown(); err != nil {
		setup.logger.Error(err)
		return err
	}
	setup.logger.Info("stopped all services")

	return nil
}

func (setup *Setup) execute(files []string) error {
	for _, file := range files {
		servicesOnFile := &Services{}
		if _, err := ReadFile(file, servicesOnFile); err != nil {
			return err
		}

		array, err := setup.load(servicesOnFile)
		if err != nil {
			return err
		}
		setup.services = append(setup.services, array...)
	}

	setup.runner = setup.NewRunner(setup.services)
	if err := setup.runner.Setup(); err != nil {
		return err
	}

	setup.logger.Info("started all services")

	if !setup.isToRunInBackground {
		setup.Wait()
	}

	return nil
}

// load recursive load services files inside every service
func (setup *Setup) load(service *Services) ([]*Services, error) {
	setup.logger.Info("loading service...")
	array := make([]*Services, 0)

	for _, file := range service.Files {
		setup.logger.Infof("loading service file %s", file)
		nextService := &Services{}
		if _, err := ReadFile(file, nextService); err != nil {
			return nil, err
		}

		setup.logger.Infof("getting next service...")
		if nextArray, err := setup.load(nextService); err != nil {
			return nil, err
		} else {
			array = append(array, nextArray...)
		}
	}

	return append(array, service), nil
}

// Wait ...
func (setup *Setup) Wait() {
	setup.logger.Info("waiting to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
