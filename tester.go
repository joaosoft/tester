package tester

import (
	"path/filepath"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
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
	config, simpleConfig, err := NewConfig()
	pm := manager.NewManager(manager.WithRunInBackground(false))

	test := &Tester{
		tests:  make(map[string]*TestFile, 0),
		config: &config.Tester,
	}

	if test.isLogExternal {
		pm.Reconfigure(manager.WithLogger(log))
	}

	if err != nil {
		log.Error(err.Error())
	} else {
		test.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Tester.Log.Level)
		log.Debugf("setting log level to %s", level)
		log.Reconfigure(logger.WithLevel(level))
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
