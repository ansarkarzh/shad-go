//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	builder := strings.Builder{}
	builder.Grow(len(input)*3)

	remaining := input
	for len(remaining) > 0 {
		r, size := utf8.DecodeLastRuneInString(remaining)
		builder.WriteRune(r)
		remaining = remaining[:len(remaining)-size]
	}
	return builder.String()
}
