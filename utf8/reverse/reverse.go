//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	var output strings.Builder

	for len(input) > 0 {
		r, size := utf8.DecodeLastRuneInString(input)
		output.WriteRune(r)
		input = input[:len(input)-size]
	}

	return output.String()
}
