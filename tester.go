package tester

import (
	"path/filepath"

	"github.com/labstack/gommon/log"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// Test ...
type Tester struct {
	tests         map[string]*TestFile
	runner        IRunner
	config        *TesterConfig
	pm            *manager.Manager
	logger        logger.ILogger
	isLogExternal bool
}

// NewGoTest ...make
func NewTester(options ...TesterOption) *Tester {
	config, simpleConfig, err := NewConfig()

	service := &Tester{
		tests:  make(map[string]*TestFile, 0),
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		logger: logger.NewLogDefault("tester", logger.WarnLevel),
		config: &config.Tester,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	if err != nil {
		log.Error(err.Error())
	} else {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Tester.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	service.Reconfigure(options...)

	return service
}

// Run ...
func (tester *Tester) Run() error {
	files, err := filepath.Glob(global[path_key].(string) + "*.json")
	if err != nil {
		return err
	}
	if err := tester.execute(files); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// RunSingle ...
func (test *Tester) RunSingle(file string) error {
	if err := test.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (test *Tester) execute(files []string) error {
	for _, file := range files {
		log.Infof("loading test file %s", file)
		testsOnFile := &TestFile{}
		if _, err := ReadFile(file, testsOnFile); err != nil {
			return err
		}
		test.tests[file] = testsOnFile
	}

	test.runner = test.NewRunner(test.tests)
	if err := test.runner.Run(); err != nil {
		log.Info("error running test files")
		return err
	}

	return nil
}
