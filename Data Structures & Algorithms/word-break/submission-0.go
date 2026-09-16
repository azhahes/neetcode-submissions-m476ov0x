func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true
	for i:=len(s)-1;i>=0;i--{
		for _, word := range wordDict{
			if len(s) >= len(word)+i && 
				s[i:i+len(word)] == word && 
				dp[i+len(word)]{
					dp[i] = true
					break
			}
		}
	}
	return dp[0]
}
