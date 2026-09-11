func productExceptSelf(nums []int) []int {
	prevProd := make([]int, len(nums))
	prevProd[0] = 1

	for i := 1; i < len(nums); i++{
		prevProd[i] = prevProd[i-1]* nums[i-1]
	}

	prev := 1
	for i := len(nums)-1; i >=0; i--{
		prevProd[i] *= prev
		prev *= nums[i]
	}
	return prevProd		
}
/*
[1,1,2,8]
[1,2,4,6]
*/