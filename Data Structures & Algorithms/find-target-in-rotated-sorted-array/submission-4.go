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
	return left
}

func search(nums []int, target int) int {
	left, right := findMin(nums), len(nums)-1
	if nums[right] < target {
		left, right =0, left-1
	}  
	for left<right {
		mid := left +(right-left)/2
		if nums[mid]>=target{
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target{
		return left
	}
	return -1
}
