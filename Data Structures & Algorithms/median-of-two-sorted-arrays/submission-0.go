func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2){
		nums1, nums2 = nums2, nums1
	}
	n1, n2 := len(nums1), len(nums2)
	t := (n1 + n2)
	h := (t+1)/2 // t=2 -> h=1; t=3 -> h=2
	l, r := -1, n1-1
	for l<=r {
		i1 := l + (r-l)/2
		i2 := h - i1 - 2

		left1, right1 := math.MinInt, math.MaxInt
		left2, right2 := math.MinInt, math.MaxInt

		if i1 >= 0{
			left1 = nums1[i1]
		}
		if i1+1 < n1{
			right1 = nums1[i1+1]
		}

		if i2 >= 0{
			left2 = nums2[i2]
		}
		if i2+1 < n2{
			right2 = nums2[i2+1]
		}

		if left1<=right2 && left2<=right1{
			if (t % 2) == 0 {
				return (float64(max(left1, left2)) + float64(min(right1, right2)))/2.0
			} else {
				return float64(max(left1, left2))
			}
		}

		if left1 > right2 {
			r = i1 - 1
		} else {
			l = i1 + 1
		} 
	}
	return -1
}
/*


[1,1,1,1,2]
[6,7,8]

[1,2,3,4,5]
[6,7,8,9,10]

[1,3,5,7,9]
[2,4,6,8,10]

[1,2,3,4,5]
[1,2,3,4,5]
*/