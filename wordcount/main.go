//go:build !solution

package main

import (
	"fmt"
	"os"
	//"path/filepath"
	s "strings"
)

func main() {
	fileNames := os.Args[1:]
	lines := make(map[string]int)
	for _, name := range fileNames {
		// path := filepath.Join(os.TempDir(), name)
		dat, err := os.ReadFile(name)
		if (err != nil) {
			panic(err)
		}

		for _, line := range s.Split(string(dat), "\n") {
			lines[line]++
		}
	}

	for line, count := range lines {
		if (count >= 2) {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
