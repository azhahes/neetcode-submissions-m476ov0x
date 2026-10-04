type MinStack struct {
	s []int
	m []int
}

func Constructor() MinStack {
	return MinStack{ []int{}, []int{} }
}

func (this *MinStack) Push(val int) {
	if this.GetMin() >= val{
		this.m = append(this.m, val)
	}
	this.s = append(this.s, val)
}

func (this *MinStack) Pop() {
	n:= len(this.s)
	if n == 0{
		return
	}
	val := this.s[n-1]
	if this.GetMin() == val{
		i:= len(this.m)
		this.m = this.m[:i-1]
	} 
	this.s = this.s[:n-1]
}

func (this *MinStack) Top() int {
	n:= len(this.s)
	return this.s[n-1]
}

func (this *MinStack) GetMin() int {
	n:= len(this.m)
	if n == 0{
		return math.MaxInt
	}
	return this.m[n-1]
}
