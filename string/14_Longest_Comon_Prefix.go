package strings

import "math"

// MyOwnSolution uses the simplest method to compare m*n times.
func MyOwnSolution14(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLength := math.MaxInt
	strLists := make([][]rune, 0, len(strs))
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
		strLists = append(strLists, []rune(str))
	}

	flag := 0
	for i := 0; i < minLength; i++ {
		tmp := strLists[0][i]
		for j := 0; j < len(strs); j++ {
			if strLists[j][i] != tmp {
				flag = i
				goto end
			}
		}
	}
end:
	if flag == 0 {
		return ""
	}
	return strs[0][:flag]
}
