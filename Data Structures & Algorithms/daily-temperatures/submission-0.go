func dailyTemperatures(temperatures []int) []int {
	 res := make([]int, len(temperatures))
	 stack := make([]int, 0)
	 stack = append(stack, 0)
	 for i := 1; i< len(temperatures);i++{
		for len(stack)>0{
			top := stack[len(stack)-1]
			if temperatures[top]>=temperatures[i]{
				break
			}
			res[top] = i-top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	 }
	 return res
}
