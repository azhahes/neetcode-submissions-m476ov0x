func largestRectangleArea(heights []int) int {
	minRight := getMinRight(heights)
	minLeft := getMinLeft(heights)
	res := 0
	for i:=0;i<len(heights);i++{
		l := (minRight[i]-minLeft[i])-1
		res = max(res, heights[i] * l)
	}
	return res
}

func getMinRight(heights []int) []int{
	n := len(heights)
	minRight := make([]int, n)
	for i := range minRight {
		minRight[i] = n
	}
	for i := n-2; i>=0 ; i-- {
		j := i+1
		for j<n && heights[j]>=heights[i]{
			j=minRight[j]
		}
		minRight[i] = j
	}
	return minRight
}

func getMinLeft(heights []int) []int{
	n := len(heights)
	minLeft := make([]int, n)
	for i := range minLeft {
		minLeft[i] = -1
	}
	for i := 1; i<n ; i++ {
		j := i-1
		for j>=0 && heights[j]>=heights[i]{
			j=minLeft[j]
		}
		minLeft[i] = j
	}
	return minLeft
}