package gotest

import (
	"path/filepath"

	"github.com/joaosoft/go-log/service"
)

// GoTest ...
type GoTest struct {
	tests  map[string]*TestFile
	runner IRunner
	config *AppConfig
}

// NewGoTest ...make
func NewGoTest(options ...GoTestOption) *GoTest {
	log.Info("starting GoTest Service")

	// load configuration file
	configApp := &AppConfig{}
	if _, err := readFile("./config/app.json", configApp); err != nil {
		log.Error(err)
	} else {
		level, _ := golog.ParseLevel(configApp.Log.Level)
		log.Debugf("setting log level to %s", level)
		WithLogLevel(level)
	}

	test := &GoTest{
		tests:  make(map[string]*TestFile, 0),
		config: configApp,
	}

	global["path"] = defaultPath

	test.Reconfigure(options...)

	return test
}

// Run ...
func (gotest *GoTest) Run() error {
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
