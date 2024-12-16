package classical_150

// maxProfit 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	minValue, ans := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		if prices[i]-minValue > ans {
			ans = prices[i] - minValue
		}
	}
	return ans
}
