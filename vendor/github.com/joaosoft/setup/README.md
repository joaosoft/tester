# setup
[![Build Status](https://travis-ci.org/joaosoft/setup.svg?branch=master)](https://travis-ci.org/joaosoft/setup) | [![codecov](https://codecov.io/gh/joaosoft/setup/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/setup) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/go-test)](https://goreportcard.com/report/github.com/joaosoft/go-test) | [![GoDoc](https://godoc.org/github.com/joaosoft/setup?status.svg)](https://godoc.org/github.com/joaosoft/setup)

A framework that allows you to setup services. At the moment it has support for web services, redis, postgres, mysql and nsq services. 
This frameworks runs all real services allowing you to validade the integration between services and your own code.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* HTTP
* SQL
* NSQ
* REDIS

## Dependency Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

>### Go
```
go get github.com/joaosoft/setup
```

## Docker
>### Start Environment 
* Redis / Postgres / MySQL / NSQ
```
make env
```

## Usage 
This example is available in the project at [setup/example](https://github.com/joaosoft/setup/tree/master/example)

```go
package main

import (
	setup setup
)

func main() {
	test := setup.NewGoSetup(
    		setup.WithPath("./examples"),
    		setup.WithRunInBackground(true))
    
    //// web
    //test.RunSingle("001_webservices.json")
    //
    //// sql
    //configSQL := &setup.SQLConfig{
    //	Driver:     "postgres",
    //	DataSource: "postgres://user:password@localhost:7001?sslmode=disable",
    //}
    //test.Reconfigure(setup.WithSQLConfiguration(configSQL))
    //test.RunSingle("002_sql.json")
    //
    //// nsq
    //configNSQ := &setup.NSQConfig{
    //	Lookupd:      "localhost:4150",
    //	RequeueDelay: 30,
    //	MaxInFlight:  5,
    //	MaxAttempts:  5,
    //}
    //test.Reconfigure(setup.WithNSQConfiguration(configNSQ))
    //test.RunSingle("003_nsq.json")
    //
    //// redis
    //configRedis := &setup.RedisConfig{
    //	Protocol: "tcp",
    //	Address:  "localhost:6379",
    //	Size:     10,
    //}
    //test.Reconfigure(setup.WithRedisConfiguration(configRedis))
    //test.RunSingle("004_redis.json")
    
    //// files
    //test.RunSingle("005_files.json")

    // all
    test.Reconfigure(
        setup.WithConfigurationFile("data/config.json"))

    test.Run()
    test.Wait()
    test.Stop()
}
```

>## Configurations

>### WebServices [ see 001_http.json ] [setup/example/001_http.json](https://github.com/joaosoft/setup/tree/master/example/001_http.json)

```javascript
{
  "http": [
    {
      "name": "hello",
      "description": "test hello",
      "host": ":8001",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "GET",
          "route": "/hello",
          "response": {
            "status": 200,
            "body": {
              "message": "Hello friend!"
            }
          }
        }
      ]
    },
    {
      "name": "goodbye",
      "description": "test goodbye",
      "host": ":8002",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "GET",
          "route": "/goodbye",
          "response": {
            "status": 200,
            "body": {
              "message": "Goodbye friend!"
            }
          }
        }
      ]
    },
    {
      "name": "something",
      "description": "testing payload of a post",
      "host": ":8003",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "POST",
          "route": "/something",
          "body": {
            "name": "joao",
            "age": 29
          },
          "response": {
            "status": 200,
            "body": {
              "message": "Goodbye friend!"
            }
          }
        }
      ]
    },
    {
      "name": "loading",
      "description": "loading the payload from a file",
      "host": ":8004",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "POST",
          "route": "/loading",
          "headers": {
            "Cookie": ["Cookie_2=value002; Cookie_1=value001"],
            "Accept-Encoding": ["gzip, deflate"],
            "Accept": ["*/*"],
            "Connection": ["keep-alive"],
            "User-Agent": ["PostmanRuntime/7.1.1"],
            "Cache-Control": ["no-cache"],
            "Content-Length": ["33"],
            "Content-Type": ["application/json"]
          },
          "cookies": [
            {
              "name": "Cookie_1",
              "value": "value001"
            },
            {
              "name": "Cookie_2",
              "value": "value002"
            }
          ],
          "file": "data/http_body_request.json",
          "response": {
            "status": 200,
            "file": "data/http_body_response.json"
          }
        }
      ]
    },
    {
      "name": "loading",
      "description": "loading the payload from a file",
      "host": ":8005",
      "routes": [
        {
          "description": "creating web mock service",
          "method": "POST",
          "route": "/loading",
          "file": "data/http_body_request.json",
          "response": {
            "status": 200,
            "file": "data/http_body_response.json"
          }
        }
      ]
    }
  ]
}
```

>### SQL [ see 002_sql.json ] [setup/example/002_sql.json](https://github.com/joaosoft/setup/tree/master/example/002_sql.json)
```javascript
{
  "sql": [
    {
      "name": "postgres",
      "description": "add users information",
      "configuration": {
        "driver": "postgres",
        "datasource": "postgres://user:password@localhost:7001?sslmode=disable"
      },
      "run": {
        "setup": [
          {
            "queries": [
              "DROP TABLE IF EXISTS USERS",
              "CREATE TABLE USERS(name varchar(255), description varchar(255))",
              "INSERT INTO USERS(name, description) VALUES('joao', 'administrator')",
              "INSERT INTO USERS(name, description) VALUES('tiago', 'user')"
            ]
          }
        ],
        "teardown": [ {
            "queries": [
              "DROP TABLE IF EXISTS USERS"
            ]
          }
        ]
      }
    },
    {
      "name": "postgres",
      "description": "add users information from files",
      "run": {
        "setup": [
            {
              "files": ["data/sql_setup_file.sql"]
            }
          ],
        "teardown": [
          {
            "files": ["data/sql_teardown_file.sql"]
          }
        ]
      }
    },
    {
      "name": "mysql",
      "description": "add clients information",
      "configuration": {
        "driver": "mysql",
        "datasource": "root:password@tcp(127.0.0.1:7002)/mysql"
      },
      "run": {
        "setup": [
          {
            "queries": [
              "DROP TABLE IF EXISTS CLIENTS",
              "CREATE TABLE CLIENTS(name varchar(255), description varchar(255))",
              "INSERT INTO CLIENTS(name, description) VALUES('joao', 'administrator')",
              "INSERT INTO CLIENTS(name, description) VALUES('tiago', 'user')"
            ]
          }
        ],
        "teardown": [
          {
            "queries": [
              "DROP TABLE IF EXISTS CLIENTS"
            ]
          }
        ]
      }
    }
  ]
}
```

>### NSQ [ see 003_nsq.json ] [setup/example/003_nsq.json](https://github.com/joaosoft/setup/tree/master/example/003_nsq.json)
```javascript
{
  "nsq": [
    {
      "name": "nsq",
      "description": "loading a script from file and from body",
      "configuration": {
        "lookupd": "localhost:4150",
        "requeue_delay": 30,
        "max_in_flight": 5,
        "max_attempts": 5
      },
      "run": {
        "setup": [
          {
            "description": "ADD PERSON ONE",
            "topic": examples,
            "message": {
              "name": "joao",
              "age": 29
            }
          },
          {
            "description": "ADD PERSON ONE",
            "topic": examples,
            "file": "data/xml_file.txt"
          }
        ],
        "teardown": []
      }
    },
    {
      "name": "nsq",
      "description": "",
      "configuration": {
        "lookupd": "localhost:4150",
        "requeue_delay": 30,
        "max_in_flight": 5,
        "max_attempts": 5
      },
      "run": {
        "setup": [
          {
            "description": "ADD PERSON TWO",
            "topic": examples,
            "message": {
              "name": "pedro",
              "age": 30
            }
          },
          {
            "description": "ADD PERSON TWO",
            "topic": examples,
            "file": "data/xml_file.txt"
          }
        ],
        "teardown": []
      }
    }
  ]
}
```

>### REDIS [ see 004_redis.json ] [setup/example/004_redis.json](https://github.com/joaosoft/setup/tree/master/example/004_redis.json)
```javascript
{
  "redis": [
    {
      "name": "redis",
      "description": "loading redis commands from file",
      "configuration": {
        "protocol": "tcp",
        "address": "localhost:6379",
        "size": 10
      },
      "run": {
        "setup": [
          {
            "files": ["data/redis_setup_file.txt"]
          }
        ],
        "teardown": [
          {
            "commands": [
              {
                "command": "DEL",
                "arguments": [
                  "id"
                ]
              },
              {
                "command": "DEL",
                "arguments": [
                  "name"
                ]
              }
            ]
          }
        ]
      }
    },
    {
      "name": "redis",
      "description": "adding by commands",
      "run": {
        "setup": [
          {
            "commands": [
              {
                "command": "APPEND",
                "arguments": [
                  "id",
                  "1"
                ]
              },
              {
                "command": "APPEND",
                "arguments": [
                  "name",
                  "JOAO RIBEIRO"
                ]
              }
            ]
          }
        ],
        "teardown": [
          {
            "commands": [
              {
                "command": "APPEND",
                "arguments": [
                  "id",
                  "2"
                ]
              },
              {
                "command": "APPEND",
                "arguments": [
                  "name",
                  "PEDRO RIBEIRO"
                ]
              }
            ]
          }
        ]
      }
    }
  ]
}
```

>### FILES [ see 005_files.json ] [setup/example/005_files.json](https://github.com/joaosoft/setup/tree/master/example/005_files.json)
```javascript
{
  "files": ["001_http.json", "002_sql.json"]
}
```

>### ALL [ see 005_all.json ] [setup/example/005_all.json](https://github.com/joaosoft/setup/tree/master/example/005_all.json)
This example have all previous setup's, just to show you that you can config them all together.

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
