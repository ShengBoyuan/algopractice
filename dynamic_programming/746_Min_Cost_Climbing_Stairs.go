package dynamic_programming

func MinCostClimbingStairs(cost []int) int {
	m, n, sum := 0, 0, 0

	for i := 1; i < len(cost); i++ {
		if cost[i-1]+m > cost[i]+n {
			sum = cost[i] + n
		} else {
			sum = cost[i-1] + m
		}

		m = n
		n = sum
	}

	return sum
}
