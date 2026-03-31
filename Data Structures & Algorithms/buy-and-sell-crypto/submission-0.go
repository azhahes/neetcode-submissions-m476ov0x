func maxProfit(prices []int) int {
	left, right := 0, 0
	maxProfit := 0
	for right < len(prices){
		maxProfit = max(maxProfit, prices[right]-prices[left])
		if prices[left] > prices[right]{
			left = right
		} else {
			right++
		}
	}
	return maxProfit
}
