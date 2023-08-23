package dynamic_programming

func BuyAndSell(nums []int) int {
	low, profit := 10000, 0

	for _, v := range nums {
		if low > v {
			low = v
		} else if v-low > profit {
			profit = v - low
		}
	}

	return profit
}
