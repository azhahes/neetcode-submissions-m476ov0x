func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}	
	res := 1
	for _, num := range nums {
		n := num
		if set[n+1]{
			continue
		}
		curr := 1
		for set[n-1] {
			curr++
			n--	
		} 
		res = max(res, curr)
	}
	return res
}
