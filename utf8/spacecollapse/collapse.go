//go:build !solution

package spacecollapse

import (
	"strings"
)

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func CollapseSpaces(input string) string {
	builder := strings.Builder{}
	builder.Grow(len(input))
	previousWasSpace := true
	for i, el := range input {
		if !isSpace(el) {
			builder.WriteRune(el)
			previousWasSpace = false
			continue
		}
		if i == 0 {
			builder.WriteRune(' ')
			continue
		}
		if !previousWasSpace {
			builder.WriteRune(' ')
		}
		previousWasSpace = true
	}
	return builder.String()
}
