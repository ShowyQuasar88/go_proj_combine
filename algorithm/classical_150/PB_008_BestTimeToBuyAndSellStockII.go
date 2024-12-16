package classical_150

// maxProfitII 买卖股票的最佳时机II
func maxProfitII(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}
