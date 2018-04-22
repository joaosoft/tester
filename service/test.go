package gotest

import (
	"path/filepath"

	"fmt"

	"github.com/joaosoft/go-log/service"
	"github.com/joaosoft/go-manager/service"
)

// GoTest ...
type GoTest struct {
	tests  map[string]*TestFile
	runner IRunner
	config *goTestConfig
	pm     *gomanager.GoManager
}

// NewGoTest ...make
func NewGoTest(options ...GoTestOption) *GoTest {
	log.Info("starting GoTest Service")
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

	test := &GoTest{
		tests:  make(map[string]*TestFile, 0),
		config: &appConfig.GoTest,
	}

	test.Reconfigure(options...)

	return test
}

// Run ...
func (gotest *GoTest) Run() error {
	files, err := filepath.Glob(global[path_key].(string) + "*.json")
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
func (gotest *GoTest) RunSingle(file string) error {
	if err := gotest.execute([]string{file}); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (gotest *GoTest) execute(files []string) error {
	for _, file := range files {
		log.Infof("loading test file %s", file)
		testsOnFile := &TestFile{}
		if _, err := readFile(file, testsOnFile); err != nil {
			return err
		}
		gotest.tests[file] = testsOnFile
	}

	gotest.runner = NewRunner(gotest.tests)
	if err := gotest.runner.Run(); err != nil {
		log.Info("error running test files")
		return err
	}

	return nil
}
