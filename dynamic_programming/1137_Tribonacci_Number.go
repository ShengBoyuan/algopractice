package dynamic_programming

func Tribonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	a, b, c, t := 0, 1, 1, 0
	for i := 3; i <= n; i++ {
		t = a + b + c
		a = b
		b = c
		c = t
	}

	return t
}
