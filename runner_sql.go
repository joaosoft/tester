package tester

import (
	"database/sql"
	"fmt"
	"github.com/joaosoft/logger"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
)

type SQLRunner struct {
	tests []SqlTest
	logger logger.ILogger
}

func (runner *Runner) NewSQLRunner(scenarioRunner *ScenarioRunner, services []SqlTest) *SQLRunner {
	return &SQLRunner{
		tests: services,
		logger: runner.logger,
	}
}

func (runner *SQLRunner) Run() error {
	for _, test := range runner.tests {
		runner.logger.Infof("running sql tester with [name: %s, description %s ]", test.Name, test.Description)

		var conn *sql.DB
		var err error
		if conn, err = test.Configuration.connect(); err != nil {
			return fmt.Errorf("failed to create sql connection")
		}

		var expected *sql.Rows
		if test.Expected.Command != nil {
			if expected, err = runner.runCommand(conn, test.Expected.Command); err != nil {
				return err
			}
		} else if test.Expected.File != nil {
			if expected, err = runner.runFile(conn, test.Expected.File); err != nil {
				return err
			}
		}
		if expected.Next() {
			var result bool
			expected.Scan(&result)
			if !result {
				return fmt.Errorf("error on sql validation. it should return 'true'")
			}
		}
	}
	return nil
}

func (runner *SQLRunner) runCommand(conn *sql.DB, command *string) (*sql.Rows, error) {
	if result, err := conn.Query(*command); err != nil {
		return nil, fmt.Errorf("failed to execute sql command [ %s ]", err)
	} else {
		return result, nil
	}
}

func (runner *SQLRunner) runFile(conn *sql.DB, file *string) (*sql.Rows, error) {
	runner.logger.Infof("executing nsq commands by file [ %s ]", *file)

	var query string
	if bytes, err := ReadFile(*file, nil); err != nil {
		return nil, fmt.Errorf("failed to read sql file [ %s ] with error [ %s ]", *file, err)
	} else {
		query = string(bytes)
	}

	return runner.runCommand(conn, &query)
}
