# go-test
[![Build Status](https://travis-ci.org/joaosoft/go-test.svg?branch=master)](https://travis-ci.org/joaosoft/go-test) | [![Code Climate](https://codeclimate.com/github/joaosoft/go-test/badges/coverage.svg)](https://codeclimate.com/github/joaosoft/go-test)

A package framework for unit and integration tests. 
## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

>### Go
```
go get github.com/joaosoft/go-test
```

## Docker
>### Start Environment 
* Redis
```
make env
```

>### Start Application
```
make start
```

## Usage 
This example is available in the project at [go-test/bin](https://github.com/joaosoft/go-test/tree/master/bin)

>### Configuration services.json
```javascript
{
  "webservices": [
    {
      "name": "hello",
      "host": ":8001",
      "routes": [
        {
          "description": "creating web test service",
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
      "host": ":8002",
      "routes": [
        {
          "description": "creating web test service",
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
    }
  ],
  "redis": [
    {
      "name": "redis",
      "configuration": {
        "protocol": "tcp",
        "addr": "redis:6379",
        "size": 10
      },
      "commands": [
        {
          "command": "DEL",
          "arguments": ["id"]
        },
        {
          "command": "DEL",
          "arguments": ["name"]
        },
        {
          "command": "APPEND", 
          "arguments": ["id", "1"]
        },
        {
          "command": "APPEND", 
          "arguments": ["name", "JOAO RIBEIRO"]
        }
      ]
    }
  ]
}
```

>### Run
```go
import "github.com/joaosoft/go-test"

func main() {
	gotest := NewGoMock(WithPath("./bin/config"), WithRunInBackground(false))
	gotest.Run()
}
```

## Run example
```
make run
```

You can see that you have created two web services:
* http://localhost:8001/hello
* http://localhost:8002/goodbye

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
