package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFibonacciEndpoint(t *testing.T) {
	url := fmt.Sprintf("http://localhost:8080/fib/nan")
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to make request to %v: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}
}
