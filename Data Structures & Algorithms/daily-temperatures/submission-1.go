func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	for i:=n-2;i>=0;i--{ 
		j := i+1
		for j<n && temperatures[j]<=temperatures[i]{
			if res[j] == 0{ // here we are checking whether we computed j already
				j=n // since not computed already, set j to max
				break // break to avoid referencing res[n] out of bound
			}
			j += res[j] // here instead of looping ot find the next high temp, we jump based on previously computed val
		}
		if j<n{
			res[i] = j-i // update the result only j is not max (n)
		}
	}
	return res
}
