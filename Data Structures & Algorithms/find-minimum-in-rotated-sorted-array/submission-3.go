func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left<right {
		mid := left + (right-left)/2
		if nums[right]>nums[mid] {
			right = mid
		} else {
			left = mid+1
		}
	}
	return nums[left]
}
