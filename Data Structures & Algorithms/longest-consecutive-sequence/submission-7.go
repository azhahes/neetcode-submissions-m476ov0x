func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}
	res := 1
	for _, n := range nums {
		if !set[n+1]{
			curr := 1
			for set[n-1] {
				curr++
				n = n-1
			}
			res = max(res, curr)
		}
	}
	return res
}
