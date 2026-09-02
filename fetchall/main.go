//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string, channel chan string) {
	resp, err := http.Get(url)
	if err != nil {
		channel <- "error"
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		channel <- "error"
		return
	}

	channel <- "success"
}

func main() {
	results := make(chan string)
	args := os.Args[1:]
	for _, url := range args {
		go fetch(url, results)
	}
	for _ = range len(args) {
		fmt.Println(<-results)
	}
	os.Exit(0)
}
