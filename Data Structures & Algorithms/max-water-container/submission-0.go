func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	maxArea := 0
	for left < right {
		maxArea = max(maxArea, min(heights[left], heights[right]) * (right-left))
		if heights[left]<heights[right]{
			left++
		} else{
			right--
		}
	}
	return maxArea
}
