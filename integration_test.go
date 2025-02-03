package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestFibonacciEndpoint(t *testing.T) {
	url := fmt.Sprintf("http://localhost:8080/fib/2")
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("Failed to make request to %v: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	expected := "1"
	if string(body) != expected {
		t.Errorf("Expected response body %s; got %s", expected, string(body))
	}
}
