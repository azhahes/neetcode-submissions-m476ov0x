func maxArea(heights []int) int {
	start, end := 0, len(heights)-1
	res := 0
	for start<end {
		res = max(res,min(heights[start], heights[end]) * (end-start))
		if heights[start]<heights[end]{
			start++
		} else{
			end--
		}
	}
	return res
}
