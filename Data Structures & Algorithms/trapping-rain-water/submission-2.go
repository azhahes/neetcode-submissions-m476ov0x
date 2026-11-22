func trap(height []int) int {
	lmax, rmax := 0,0
	start, end := 0, len(height)-1
	res := 0 
	for start<end {
		if height[start]<height[end] {
			trapped := lmax - height[start]
			if trapped>0{
				res += trapped
			} else {
				lmax = height[start]	
			}
			start++
		} else {
			trapped := rmax - height[end]
			if trapped>0{
				res += trapped
			} else {
				rmax = height[end]	
			}
			end--
		}
	}
	return res
}
