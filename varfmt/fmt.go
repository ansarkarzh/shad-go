//go:build !solution

package varfmt

import (
	"fmt"
	"strconv"
	"strings"
)

func Sprintf(format string, args ...interface{}) string {
	resB := strings.Builder{}
	resB.Grow(len(format))
	formattedArgs := make(map[int]string)
	left := 0
	right := 0
	inBraces := false
	braceIndex := 0
	for i, r := range(format) {
		if r == '{' {
			left = i
			inBraces = true
			continue
		}
		if r == '}' {
			right = i
			inBraces = false

			numberS := format[left + 1: right]
			argIndex := 0
			if len(numberS) == 0 {
				argIndex = braceIndex
			} else {
				number, err := strconv.Atoi(numberS)
				if err != nil {
					panic(1)
				}
				argIndex = number
			}
			formatted, ok := formattedArgs[argIndex]
			if ok == false {
				formatted = fmt.Sprint(args[argIndex])
				formattedArgs[argIndex] = formatted
			}
			resB.WriteString(formatted)
			braceIndex++
			continue
		}
		if !inBraces {
			resB.WriteRune(r)
		}
	}

	return resB.String()
}
