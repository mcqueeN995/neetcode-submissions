func maxProfit(prices []int) int {
	l, res := 0, 0

	for r := 1; r < len(prices); r++ {
		if prices[r] < prices[l] {
			l = r
		}

		if res < prices[r] - prices[l] {
			res = prices[r] - prices[l]
		}
	}
	return res
}
