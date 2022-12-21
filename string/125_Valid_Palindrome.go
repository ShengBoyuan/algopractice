package strings

import (
	"strings"
	"unicode"
)

// MyOwnSolution converts all elements to lower and removes non
// alphanumeric ones; then use two pointers to read forward and
// backward.
func MyOwnSolution125(s string) bool {
	words := make([]rune, 0, len(s))
	for _, t := range strings.ToLower(s) {
		if (t >= 'a' && t <= 'z') || (t >= '0' && t <= '9') {
			words = append(words, t)
		}
	}

	j := len(words) - 1
	for i := 0; i <= j; i++ {
		if words[i] != words[j] {
			return false
		}
		j--
	}

	return true
}

// StdLib uses standard libs (unicode, strings) to handle input
// string, the process is almost the same as mine.
func StdLib(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// PS: 1. Alphanumeric characters include letters and numbers;
//	   2. string.Map(f, s) seems very cool.
