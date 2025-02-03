.PHONY: run test/unit test/integration

run:
	go build -o server
	./server

test/unit:
	go test -v -run TestFibonacciUnit

test/integration:
	go test -v -run TestFibonacciEndpoint