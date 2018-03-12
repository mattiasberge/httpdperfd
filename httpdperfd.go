package main

import (
	"bytes"
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
	if len(weight) < 1 {
		difficulty := 1000
		return difficulty
	}
	difficulty, err := strconv.Atoi(weight[0])
	if err != nil {
		difficulty := 1000
		return difficulty
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

// generate blob of <size> bytes
func generate_body(rbb []string) bytes.Buffer {
	var buffer bytes.Buffer
	if len(rbb) < 1 {
		return buffer
	}
	rbb_size, err := strconv.Atoi(rbb[0])
	if err != nil {
		return buffer // empty buffer
	}
	for i := 0; i < rbb_size; i++ {
		buffer.WriteString("1")
	}
	return buffer
}

// handle all http calls
func slash(w http.ResponseWriter, r *http.Request) {
	weight, _ := r.URL.Query()["weight"]
	rbb, _ := r.URL.Query()["response_body_bytes"]
	var buffer bytes.Buffer

	response_body_bytes := generate_body(rbb)

	difficulty := set_difficulty(weight)
	elapsed := process_checksums(difficulty)

	msg := fmt.Sprintf("%s%s, difficulty: %d, elapsed: %v\n", r.Host, r.URL.Path, difficulty, elapsed)
	buffer.WriteString(msg)
	buffer.WriteString(response_body_bytes.String())

	response_body := buffer.String()

	io.WriteString(w, response_body)

	if len(os.Getenv("LOG")) < 1 {
		return
	}

	fmt.Printf(msg)
	return
}

// main
func main() {
	if len(os.Getenv("LOG")) > 0 {
		fmt.Println("Starting server with Logging enabled")
	}
	http.HandleFunc("/", slash)
	http.ListenAndServe(":8000", nil)
}
