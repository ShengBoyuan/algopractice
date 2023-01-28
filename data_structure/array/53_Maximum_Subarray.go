package array

// Don't over complicate things.
func MaxSubarray(nums []int) int {
	max, sum := nums[0], nums[0]

	for _, t := range nums[1:] {
		if sum < 0 {
			sum = t // if sum < 0, it means a totally new subarray begins.
		} else {
			sum += t
		}

		if max < sum {
			max = sum // record the maximum every single time.
		}
	}

	return max
}
