func maxSlidingWindow(nums []int, k int) []int {
   result := make([]int, 0)
   l, r := 0, 0
   q := make([]int, 0)
   for r < len(nums){
	for len(q) > 0 && nums[q[len(q)-1]] < nums[r]{
		q = q[:len(q)-1]
	}
	q = append(q, r)
	if l > q[0]{
		q = q[1:]
	}
	if r+1 >= k{
		result = append(result, nums[q[0]])
		l++
	}
	r++
   } 
   return result
}
