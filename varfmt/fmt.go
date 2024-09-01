package varfmt

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Sprintf(format string, args ...interface{}) string {
	var sb strings.Builder
	sb.Grow(len(format))

	pos := 0

	for len(format) > 0 {
		r, size, _ := nextRune(format)
		format = format[size:]

		if r == '{' {
			closingBracePos := strings.Index(format, "}")
			if closingBracePos == -1 {
				sb.WriteString(format)
				break
			}

			argIndexStr := format[:closingBracePos]
			format = format[closingBracePos+1:]

			argIndex := pos
			if argIndexStr != "" {
				_, err := fmt.Sscan(argIndexStr, &argIndex)
				if err != nil {
					sb.WriteString("{" + argIndexStr + "}")
					continue
				}
			}

			if argIndex >= 0 && argIndex < len(args) {
				argStr := fmt.Sprint(args[argIndex])
				sb.WriteString(argStr)
			}

			pos++
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func nextRune(s string) (rune, int, error) {
	if len(s) == 0 {
		return 0, 0, fmt.Errorf("empty string")
	}
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size == 1 {
		return 0, 0, fmt.Errorf("invalid utf8 encoding")
	}
	return r, size, nil
}
