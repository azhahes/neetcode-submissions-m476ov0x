func trap(height []int) int {
	l, r:= 0, len(height)-1
	lmax, rmax := 0,0
	trapped := 0
	for l<r {
		if height[l]<height[r]{
			lmax = max(lmax, height[l])
			trapped += lmax - height[l]
			l++
		} else {
			rmax = max(rmax, height[r])
			trapped += rmax - height[r]
			r--
		}
	}
	return trapped
}
