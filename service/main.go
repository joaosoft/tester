package gotest

import (
	"path/filepath"

	logger "github.com/joaosoft/go-log/service"
)

// GoTest ...
type GoTest struct {
	testFiles map[string]*TestFile
	runner    IRunner
}

// NewGoTest ...make
func NewGoTest(options ...GoTestOption) *GoTest {
	log.Info("starting GoTest Service")
	mock := &GoTest{
		testFiles: make(map[string]*TestFile, 0),
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
		gotest.testFiles[file] = testsOnFile
	}

	gotest.runner = NewRunner(gotest.testFiles)
	if err := gotest.runner.Run(); err != nil {
		log.Info("error running testFiles")
		return err
	}

	return nil
}
