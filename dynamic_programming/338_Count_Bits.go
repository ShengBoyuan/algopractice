package dynamic_programming

// 最高有效位. bits[i] = bits[i-high] + 1
func CountBits1(n int) []int {
	bits := make([]int, n+1)
	high := 0

	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			high = i
		}
		bits[i] = bits[i-high] + 1
	}

	return bits
}

// 最低有效位. bits[i] = bits[i>>1] + i&1
func CountBits2(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

// 最低设置位. bits[i] = bits[i&(i-1)] + 1
func CountBits3(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}
