package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
	Difficulty (num hashes) is weight=<int>, or defaults to 1000
	Input: ?weight=<value>
	Output: difficulty int
*/
func set_difficulty(weight []string) int {
  difficulty := 1000 // default
	if len(weight) < 1 {
		return difficulty
	}
	difficulty, err := strconv.Atoi(weight[0])
	if err != nil {
		difficulty = 1000
	}
	return difficulty
}

// Waste cpu; Iterates difficulty num of checksumming
func process_checksums(iterations int) time.Duration {
	start := time.Now()
	h := sha1.New()
	for i := 0; i < iterations; i++ {
		io.WriteString(h, strconv.Itoa(i))
	}
	elapsed := time.Since(start)
	return elapsed
}

// Write zeroes to SOMETHING
func write_body(rbb []string, w http.ResponseWriter) {
	if len(rbb) < 1 {
		return
	}
	rbb_size, err := strconv.Atoi(rbb[0])
	if err != nil {
		return
	}
	for i := 0; i < rbb_size; i++ {
		io.WriteString(w, "0")
	}
}

// handle all http calls
func slash(w http.ResponseWriter, r *http.Request) {
	weight, _ := r.URL.Query()["weight"]
	rbb, _ := r.URL.Query()["response_body_bytes"]
	difficulty := set_difficulty(weight)

	elapsed := process_checksums(difficulty)
  write_body(rbb, w)

	msg := fmt.Sprintf("%s%s, difficulty: %d, elapsed: %v\n", r.Host, r.URL.Path, difficulty, elapsed)
	io.WriteString(w, msg)

	if len(os.Getenv("LOG")) > 1 {
	  fmt.Printf(msg)
	}
}

// main
func main() {
  fmt.Println("Starting httpdperfd server")
	if len(os.Getenv("LOG")) > 0 {
		fmt.Printf("Logging enabled, LOG=%s\n", os.Getenv("LOG"))
	}
	http.HandleFunc("/", slash)
	http.ListenAndServe(":8000", nil)
}
