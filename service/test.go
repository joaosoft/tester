package gotest

import (
	"path/filepath"

	"fmt"

	"github.com/joaosoft/go-log/service"
	"github.com/joaosoft/go-manager/service"
)

// Test ...
type Test struct {
	tests         map[string]*TestFile
	runner        IRunner
	config        *TestConfig
	pm            *gomanager.Manager
	isLogExternal bool
}

// NewGoTest ...make
func NewGoTest(options ...TestOption) *Test {
	log.Info("starting Test Service")
	pm := gomanager.NewManager(gomanager.WithRunInBackground(false))

	// load configuration file
	appConfig := &appConfig{}
	if simpleConfig, err := gomanager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", getEnv()), appConfig); err != nil {
		log.Error(err.Error())
	} else {
		pm.AddConfig("config_app", simpleConfig)
		level, _ := golog.ParseLevel(appConfig.GoTest.Log.Level)
		log.Debugf("setting log level to %s", level)
		WithLogLevel(level)
	}

	test := &Test{
		tests:  make(map[string]*TestFile, 0),
		config: &appConfig.GoTest,
	}

	test.Reconfigure(options...)

	if test.isLogExternal {
		pm.Reconfigure(gomanager.WithLogger(log))
	}

	return test
}

// Run ...
func (test *Test) Run() error {
	files, err := filepath.Glob(global[path_key].(string) + "*.json")
	if err != nil {
		return err
	}
	if err := test.execute(files); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// RunSingle ...
func (test *Test) RunSingle(file string) error {
	if err := test.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (test *Test) execute(files []string) error {
	for _, file := range files {
		log.Infof("loading test file %s", file)
		testsOnFile := &TestFile{}
		if _, err := readFile(file, testsOnFile); err != nil {
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
