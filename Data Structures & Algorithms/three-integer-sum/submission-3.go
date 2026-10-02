func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i:=0;i<len(nums);i++{
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		target := -nums[i]
		start, end := i+1, len(nums)-1
		for start<end {
			currSum := nums[start] + nums[end]
			if currSum == target {
				res = append(res, []int{nums[i], nums[start], nums[end]})
				j:=start+1
				for j<end && nums[j-1]==nums[j]{
					j++
				}
				start = j
			} else if currSum > target {
				end--
			} else {
				start++
			}
		}
	}
	return res
}


/*
[-1,0,1,2,-1,-4]
[-4,-1,-1,0,1,2]
*/