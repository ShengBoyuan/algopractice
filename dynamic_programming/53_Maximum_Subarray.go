package dynamic_programming

func MaxSubarray(nums []int) int {
	max, sum := nums[0], nums[0]

	for _, t := range nums {
		if sum < 0 {
			sum = t
		} else {
			sum += t
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
