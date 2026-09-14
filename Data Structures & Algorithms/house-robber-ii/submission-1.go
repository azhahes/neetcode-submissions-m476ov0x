func rob(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
   return max(helper(nums[:len(nums)-1]), 
		   	   helper(nums[1:])) 
}

func helper(nums []int) int {
   rob1, rob2 := 0,0
   for _,n := range nums {
	rob1, rob2 = rob2, max(rob1+n, rob2)
   } 
   return rob2
}
