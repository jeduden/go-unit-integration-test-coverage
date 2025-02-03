package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "fib" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	result := fibonacci(n)
	fmt.Fprintf(w, "%d", result)
}

func main() {
	http.HandleFunc("/fib/", fibHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
