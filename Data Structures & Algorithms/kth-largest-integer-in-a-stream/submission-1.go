type IntHeap []int

func(h IntHeap)Len()int{return len(h)}
func(h IntHeap)Less(i,j int)bool{return h[i]<h[j]}
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


type KthLargest struct {
    h *IntHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
	m := IntHeap(nums)
	heap.Init(&m)
	for m.Len() > k{
		heap.Pop(&m)
	}
   return KthLargest{&m, k} 
}


func (this *KthLargest) Add(val int) int {
   heap.Push(this.h, val)
   if this.h.Len() > this.k{
	heap.Pop(this.h)
   }
   return (*this.h)[0]
}
