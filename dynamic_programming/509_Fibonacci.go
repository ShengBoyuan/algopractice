package dynamic_programming

func Fib(n int) int {
	if n == 0 {
		return 0
	}

	a, b, c := 1, 0, 0
	for i := 1; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
