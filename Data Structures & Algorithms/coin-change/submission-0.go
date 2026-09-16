func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount+1
	}
	dp[0] = 0
	for t :=1; t<=amount; t++ {
		for _, coin := range coins{
			if t-coin>=0{
				dp[t] = min(dp[t], 1+dp[t-coin])
			}
		}
	}	
	if dp[amount] == amount+1{
		return -1
	}
	return dp[amount]
}
