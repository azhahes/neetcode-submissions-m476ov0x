
func climbStairs(n int) int {
var dp = make([]int, n+1)
var dfs func (int) int
dfs = func (n int) int {
	if n<=2{
		return n
	}
	if dp[n]>0{
		return dp[n]
	}
	dp[n] = dfs(n-1) + dfs(n-2)
	return dp[n]
}
return dfs(n)
}