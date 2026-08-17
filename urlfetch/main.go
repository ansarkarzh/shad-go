//go:build !solution

package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
