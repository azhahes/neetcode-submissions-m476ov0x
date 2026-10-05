func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([][2]int, 0)
	for i, h := range heights{
		start := i
		for len(stack)>0 && stack[len(stack)-1][1] > h{
			val := stack[len(stack)-1]
			maxArea = max(maxArea, val[1] * (i-val[0]))
			stack = stack[:len(stack)-1]
			start = val[0]
		}
		stack = append(stack, [2]int{start, h})
	}
	for len(stack)>0{
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		maxArea = max(maxArea, val[1] * (len(heights)-val[0]))
	}
	return maxArea
}
