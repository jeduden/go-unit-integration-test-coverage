.PHONY: run test/unit test/integration

run:
	go build -o server
	./server

test/unit:
	go test -v -run TestFibonacci