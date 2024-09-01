//go:build !solution

package spacecollapse

func CollapseSpaces(input string) string {
	var result []rune
	var prevIsSpace bool

	for _, r := range input {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			if !prevIsSpace {
				result = append(result, ' ')
				prevIsSpace = true
			}
		} else {
			result = append(result, r)
			prevIsSpace = false
		}
	}

	return string(result)
}
