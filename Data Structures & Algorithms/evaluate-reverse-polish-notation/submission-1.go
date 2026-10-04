func evalRPN(tokens []string) int {
	exp:= map[string]func(i,j int)int{ 
		"+": func(i, j int)int{return i+j}, 
		"-": func(i, j int)int{return i-j}, 
		"*": func(i, j int)int{return i*j}, 
		"/": func(i, j int)int{return i/j},
		}
	stack := make([]int, 0)
	for _, t := range tokens {
		f, ok := exp[t]
		if !ok{
			v, _ := strconv.Atoi(t)
			stack = append(stack, v)
		} else {
			n := len(stack)
			l, r := stack[n-2], stack[n-1]
			stack = stack[:n-2]
			stack = append(stack, f(l, r))
		}
	}
	return stack[len(stack)-1]
}
