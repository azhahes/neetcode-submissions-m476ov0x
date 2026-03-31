func hasDuplicate(nums []int) bool {
    set := map[int]bool{}
	for _, n := range nums {
		if set[n]{
			return true
		}
		set[n] = true
	}
	return false
}
