.PHONY: run test/unit test/integration

clean:
	rm -rf coverage

# Run unit tests

test/unit:
	go test -v -run TestFibonacciUnit

# Measure coverage

test/unit/coverage/1:
	mkdir -p coverage/1
	go test -coverprofile=coverage/1/report.txt -v -run TestFibonacciUnit

coverage/html/1:
	go tool cover -html=coverage/1/report.txt

# Run integration test
run/1:
	go build -o server
	./server

test/integration/1:
	go test -v -run TestFibonacciEndpoint

# Measure coverage integration test
run/2:
	mkdir -p coverage/2
	go build -cover -o server
	GOCOVERDIR=$(PWD)/coverage/2 ./server

test/integration/2:
	go test -v -run TestFibonacciEndpoint


# Measure coverage integration test attempt #2
run/3:
	mkdir -p coverage/3
	go build -cover -o server
	GOCOVERDIR=$(PWD)/coverage/3 ./server

test/integration/3:
	go test -v -run TestFibonacciEndpoint
	curl http://localhost:8080/stop

coverage/percent/3:
	go tool covdata percent -i=coverage/3

#coverage/html/3:
# doesnt work

coverage/convert/3:
	go tool covdata textfmt -i=coverage/3 -o coverage/3/report.txt

coverage/html/3:
	go tool cover -html=coverage/3/report.txt

## How to merge unit test coverage?

test/unit/coverage/3:
	mkdir -p coverage/3
	go test -cover -v -run TestFibonacciUnit -args -test.gocoverdir="$(PWD)/coverage/3"

# run again
# - make coverage/convert/3 coverage/html/3

## 😎 Report

## Stopping server => 😭

run/4:
	mkdir -p coverage/4
	go build -cover -o server
	GOCOVERDIR=$(PWD)/coverage/4 ./server

test/integration/4:
	mkdir -p coverage/4
	go test -v -run TestFibonacciEndpoint
	curl http://localhost:8080/coverage

# Doesnt work, but this does

run/5:
	mkdir -p coverage/5
	go build -covermode=atomic -cover -o server
	GOCOVERDIR=$(PWD)/coverage/5 ./server

test/integration/5:
	mkdir -p coverage/5
	go test -v -run TestFibonacciEndpoint
	curl http://localhost:8080/coverage

coverage/html/5:
	go tool covdata textfmt -i=coverage/5 -o coverage/5/report.txt
	go tool cover -html=coverage/5/report.txt

# Some issues are still there:
# - Reset counters between test runs
# - How to get coverage files when the server is running in a docker container on linux