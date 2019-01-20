package tester

import (
	"fmt"
	"strings"

	redis "github.com/go-redis/redis"
)

type RedisRunner struct {
	tests []RedisTest
}

func NewRedisRunner(scenarioRunner *ScenarioRunner, services []RedisTest) *RedisRunner {
	return &RedisRunner{
		tests: services,
	}
}

func (runner *RedisRunner) Run() error {
	for _, test := range runner.tests {
		log.Infof("running redis tester with [name: %s, description %s ]", test.Name, test.Description)

		var conn *redis.Client
		var err error
		if conn, err = test.Configuration.connect(); err != nil {
			return fmt.Errorf("failed to create redis connection")
		}

		if test.Expected.Command != nil {
			if err := runner.runCommand(conn, test.Expected.Command, test.Expected.Arguments); err != nil {
				return err
			}
		} else if test.Expected.File != nil {
			if err := runner.runFile(conn, test.Expected.File); err != nil {
				return err
			}
		}
	}
	return nil
}

func (runner *RedisRunner) runCommand(conn *redis.Client, command *string, arguments []string) error {
	log.Infof("executing redis command [ %s ] arguments [ %s ]", command, arguments)
	if err := conn.Do(*command, arguments).Err(); err != nil {
		return err
	}
	return nil
}

func (runner *RedisRunner) runFile(conn *redis.Client, file *string) error {
	log.Infof("executing redis commands by file [ %s ]", *file)

	if lines, err := ReadFileLines(*file); err != nil {
		for _, line := range lines {
			command := strings.SplitN(line, " ", 1)
			arguments := strings.Split(command[1], " ")
			log.Infof("executing redis command [ %s ] arguments [ %s ]", command[0], arguments)
			return runner.runCommand(conn, &command[0], arguments)
		}
	}
	return nil
}
