type point struct{
	p [2]int
	d float64
}

type IntHeap []*point

func(h IntHeap)Len()int{return len(h)}
func(h IntHeap)Less(i,j int)bool{return h[i].d<h[j].d}
func(h IntHeap)Swap(i,j int){h[i], h[j] = h[j], h[i]}

func(h *IntHeap)Push(v any){
	*h = append(*h, v.(*point))
}

func(h *IntHeap)Pop() any{
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func kClosest(points [][]int, k int) [][]int {
	h := IntHeap{}
	heap.Init(&h)
	for _, v := range points {
		disFromOrigin := math.Sqrt(math.Pow(float64(v[0]),2)+math.Pow(float64(v[1]),2))
		p := &point{[2]int{v[0],v[1]}, disFromOrigin}
		heap.Push(&h, p)
	}
	res := make([][]int, 0)
	for range k {
		v := heap.Pop(&h).(*point)
		res = append(res, []int{v.p[0], v.p[1]})
	}
	return res
}
