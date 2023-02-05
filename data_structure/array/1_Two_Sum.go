package array

// Use Map to check whether something exists.
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		if idx, ok := numMap[target-num]; ok {
			return []int{idx, i}
		}
		numMap[num] = i // "Two" Sum, so you have 2 chances to seize the []int.
	}

	return []int{0, 0}
}
