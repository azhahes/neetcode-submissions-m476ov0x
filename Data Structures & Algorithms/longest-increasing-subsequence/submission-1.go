import (
	"slices"
)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
   dp := make([]int, len(nums))
   for i:= range dp {
	dp[i] = 1
   }
   for i:=1; i<len(dp); i++{
	for j:=i-1; j>=0; j--{
		if nums[j]<nums[i]{
			dp[i] = max(dp[i], 1+dp[j])
		}
	}
   } 
   return slices.Max(dp)
}
