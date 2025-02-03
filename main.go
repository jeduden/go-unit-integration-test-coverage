package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func stopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
	go func() {
		time.Sleep(100 * time.Millisecond) // Brief delay to ensure response is sent
		os.Exit(0)
	}()
}

func coverageHandler(w http.ResponseWriter, r *http.Request) {
	coverDir, _ := os.LookupEnv("GOCOVERDIR")
	if err := writeCoverageFiles(coverDir); err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, "OK")
	}
}

func main() {
	http.HandleFunc("/fib/", fibHandler)
	http.HandleFunc("/stop", stopHandler)
	http.HandleFunc("/coverage", coverageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
