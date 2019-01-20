package tester

import (
	"path/filepath"

	"fmt"

	logger "github.com/joaosoft/logger"
	manager "github.com/joaosoft/manager"
)

// Test ...
type Tester struct {
	tests         map[string]*TestFile
	runner        IRunner
	config        *TesterConfig
	pm            *manager.Manager
	isLogExternal bool
}

// NewGoTest ...make
func NewTester(options ...TesterOption) *Tester {
	log.Info("starting Test Service")
	pm := manager.NewManager(manager.WithRunInBackground(false))

	test := &Tester{
		tests: make(map[string]*TestFile, 0),
	}

	if test.isLogExternal {
		pm.Reconfigure(manager.WithLogger(log))
	}

	// load configuration file
	appConfig := &AppConfig{}
	if simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		log.Error(err.Error())
	} else if appConfig.Tester != nil {
		pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(appConfig.Tester.Log.Level)
		log.Debugf("setting log level to %s", level)
		log.Reconfigure(logger.WithLevel(level))
		test.config = appConfig.Tester
	}

	test.Reconfigure(options...)

	return test
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

	test.runner = NewRunner(test.tests)
	if err := test.runner.Run(); err != nil {
		log.Info("error running test files")
		return err
	}

	return nil
}
