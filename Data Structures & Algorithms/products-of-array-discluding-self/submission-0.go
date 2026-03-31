func productExceptSelf(nums []int) []int {
	n := len(nums)
	preProd := make([]int, n) 
	preProd[0] = 1
	for i:=1;i<n; i++{
		preProd[i] = nums[i-1] * preProd[i-1]
	}
	curr := 1
	for i := n-1; i>=0; i--{
		preProd[i] *= curr
		curr *= nums[i]
	}
	return preProd

}
