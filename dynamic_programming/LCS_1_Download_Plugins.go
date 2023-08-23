package dynamic_programming

func DownloadPlugins(n int) int {
	per, ans := 1, 0

	for n > 0 {
		if n <= per {
			ans++
			n -= per
		} else {
			ans++
			per *= 2
		}
	}

	return ans
}
