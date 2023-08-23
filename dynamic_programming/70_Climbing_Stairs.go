package dynamic_programming

func ClimbingStairs(n int) int {
	if n == 1 {
		return 1
	}

	a, b, c := 1, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}
