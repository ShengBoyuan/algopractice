package dynamic_programming

func Massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp0, dp1, tdp0, tdp1 := 0, nums[0], 0, 0
	for i := 1; i < n; i++ {
		if dp0 > dp1 {
			tdp0 = dp0
		} else {
			tdp0 = dp1
		}

		tdp1 = dp0 + nums[i]
		dp0, dp1 = tdp0, tdp1
	}

	if dp0 > dp1 {
		return dp0
	}
	return dp1
}
