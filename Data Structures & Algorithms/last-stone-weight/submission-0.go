type IntHeap []int

func(h IntHeap)Len()int{return len(h)}
func(h IntHeap)Less(i,j int)bool{return h[i]>h[j]}
func(h IntHeap)Swap(i,j int){h[i], h[j] = h[j], h[i]}

func(h *IntHeap)Push(v any){
	*h = append(*h, v.(int))
}

func(h *IntHeap)Pop() any{
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func lastStoneWeight(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)
	for h.Len() > 0 {
		if h.Len()==1{
			return h[0]
		}
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if x > y {
			heap.Push(&h, x-y)
		}
	} 
	return 0
}
