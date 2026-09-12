func maxProfit(prices []int) int {
	buy, sell := 0,0
	maxProfit := 0
	for sell < len(prices){
		maxProfit = max(maxProfit, prices[sell]-prices[buy])
		if prices[buy]>prices[sell]{
			buy++
		} else {
			sell++
		}
	}
	return maxProfit
}
