package array

func FindMaxConsecutiveOnes(nums []int) int {
	tmp, max := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			if tmp > max {
				max = tmp
			}
			tmp = 0
		}
	}

	if tmp > max {
		return tmp
	}
	return max
}
