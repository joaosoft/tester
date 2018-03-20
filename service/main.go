package gotest

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	logger "github.com/joaosoft/go-log/service"
)

// GoMock ...
type GoMock struct {
	services   []*Services
	runner     IRunner
	background bool
}

// NewGoMock ...make
func NewGoMock(options ...GoMockOption) *GoMock {
	log.Info("starting GoMock Service")
	mock := &GoMock{
		background: background,
		services:   make([]*Services, 0),
	}

	global["path"] = defaultPath

	// load configuration file
	app := &App{}
	if _, err := readFile("config/app.json", app); err != nil {
		log.Error(err)
	} else {
		level, _ := logger.ParseLevel(app.Log.Level)
		log.Debugf("setting log level to %s", level)
		WithLogLevel(level)
	}

	mock.Reconfigure(options...)

	return mock
}

// Run ...
func (gotest *GoMock) Run() error {
	files, err := filepath.Glob(global["path"].(string) + "*.json")
	if err != nil {
		return err
	}
	if err := gotest.execute(files); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// RunSingle ...
func (gotest *GoMock) RunSingle(file string) error {
	if err := gotest.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Stop ...
func (gotest *GoMock) Stop() error {
	if err := gotest.runner.Teardown(); err != nil {
		log.Error(err)
		return err
	}
	log.Info("stopped all services")

	return nil
}

func (gotest *GoMock) execute(files []string) error {
	for _, file := range files {
		servicesOnFile := &Services{}
		if _, err := readFile(file, servicesOnFile); err != nil {
			return err
		}
		gotest.services = append(gotest.services, servicesOnFile)
	}

	gotest.runner = NewRunner(gotest.services)
	if err := gotest.runner.Setup(); err != nil {
		return err
	}

	log.Info("started all services")

	if !gotest.background {
		gotest.Wait()
	}

	return nil
}

// Wait ...
func (gotest *GoMock) Wait() {
	log.Info("waiting to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
