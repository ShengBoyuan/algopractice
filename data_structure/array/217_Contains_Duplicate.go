package array

// Easy, nothing to say.
func ContainsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, t := range nums {
		_, ok := set[t]
		if ok {
			return true
		} else {
			set[t] = struct{}{}
		}
	}
	return false
}
