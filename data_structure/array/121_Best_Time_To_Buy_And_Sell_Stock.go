package array

func BuyAndSell(prices []int) int {
	low, profit := 10000, 0

	for _, v := range prices {
		if v <= low {
			low = v
		} else if v-low >= profit {
			profit = v - low
		}
	}

	return profit
}
