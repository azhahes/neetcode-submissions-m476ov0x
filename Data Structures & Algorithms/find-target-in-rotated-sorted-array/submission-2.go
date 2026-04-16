func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start<end {
		mid := start + ((end-start)/2)
		if nums[mid] < nums[end]{
			end = mid
		} else {
			start = mid+1
		}
	}
	n := len(nums)
	if target <= nums[n-1] {
		end = n-1
	} else {
		end = start-1
		start = 0
	}
	for start<end {
		mid := start + ((end-start)/2)
		if nums[mid] >= target{
			end = mid
		} else {
			start = mid+1
		}
	}
	if target == nums[start]{
		return start
	}
	return -1
}