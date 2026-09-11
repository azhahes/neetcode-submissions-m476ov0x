func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make(map[[3]int]bool)
	for i, num := range nums {
		target := -num
		dict := make(map[int]bool)
		j := i+1
		for j<len(nums){
			if dict[target-nums[j]]{
				result[[3]int{-target, nums[j], target-nums[j]}]=true
			}
			dict[nums[j]] = true
			j++
		}
	}
	res := make([][]int, 0)
	for k,_ := range result {
		res = append(res, k[:])
	}
	return res
}
