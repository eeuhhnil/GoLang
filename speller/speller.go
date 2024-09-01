//go:build !solution

package speller

var ones = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var tens = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	}

	if n < 0 {
		return "minus " + Spell(-n)
	}

	result := ""

	if n >= 1e9 {
		result += Spell(n/1e9) + " billion"
		n %= 1e9
	}

	if n >= 1e6 {
		if result != "" {
			result += " "
		}
		result += Spell(n/1e6) + " million"
		n %= 1e6
	}

	if n >= 1e3 {
		if result != "" {
			result += " "
		}
		result += Spell(n/1e3) + " thousand"
		n %= 1e3
	}

	if n >= 100 {
		if result != "" {
			result += " "
		}
		result += Spell(n/100) + " hundred"
		n %= 100
	}

	if n > 0 {
		if result != "" {
			result += " "
		}
		if n < 20 {
			result += ones[n]
		} else {
			result += tens[n/10]
			if n%10 > 0 {
				result += "-" + ones[n%10]
			}
		}
	}

	return result
}
