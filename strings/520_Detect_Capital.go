package strings

import "strings"

// MyOwnSolution doesn't handle three situations properly, it just awkwardly handles these
// possibilities seperately. But I think the "decreasing" idea is quite a shining point of
// this puzzle, so I keep my own solution, though it's very clumsy.
func MyOwnSolution520(word string) bool {
	if word == "" {
		return true
	}

	words := []rune(word)
	bigs := 0
	for i := 1; i < len(words); i++ {
		bigs = bigs + isCapital(words[i])
		if isCapital(words[i])-isCapital(words[i-1]) > 0 {
			return false
		}
	}

	if bigs == 0 || bigs == len(words)-1 {
		return true
	}
	return false
}

func isCapital(letter rune) int {
	if letter >= 'A' && letter <= 'Z' {
		return 1
	}
	return 0
}

// OneLineSolution is the fastest and neatest solution of my opinion, it gracefully seperate all "true"
// siturations in one line. It shows the best use of std libs and slice operations.
func OneLineSolution(word string) bool {
	return strings.ToUpper(word) == word || strings.ToLower(word) == word || strings.ToLower(word[1:]) == word[1:]
}

// Negation (I name this solution this) take the negative proposition of this problem. It defines when
// func returns "false", and "true" otherwise.
func Negation(word string) bool {
	usedNonCapital := false
	usedCapital := true
	for i := 0; i < len(word); i++ {
		if rune(word[i]) >= 'a' {
			usedNonCapital = true
		} else {
			if i > 0 {
				usedCapital = true
			}
		}
		if usedNonCapital && usedCapital {
			return false
		}
	}
	return true
}

// PS: 1. the 1st "i" of for range begins with "0";
// 	   2. boolean types can't perform bit operations.
