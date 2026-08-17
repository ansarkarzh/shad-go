//go:build !solution

package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func fetch(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	return err
}

func main() {
	exitCode := 0
	for _, url := range os.Args[1:] {
		if err := fetch(url, os.Stdout); err != nil {
			fmt.Printf("ERERE")
			exitCode = 1
		}
	}
	os.Exit(exitCode)
}
