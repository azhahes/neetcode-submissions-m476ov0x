type Item struct{
	val int
	priority int
}

type PriorityQueue []*Item

func(pq PriorityQueue) Len()int {return len(pq)}
func(pq PriorityQueue) Less(i,j int)bool{return pq[i].priority > pq[j].priority}
func(pq PriorityQueue) Swap(i,j int){pq[i], pq[j] = pq[j], pq[i] }

func(pq *PriorityQueue) Push(a any){
	*pq = append(*pq, a.(*Item))
 }

func(pq *PriorityQueue) Pop() any{ 
	old := *pq
	n := len(old)
	val := old[n-1]
	*pq = old[:n-1]
	return val
}

func maxSlidingWindow(nums []int, k int) []int {
   pq := make(PriorityQueue, 0)
   heap.Init(&pq)
   l, r := 0, 0
   for ;r<k ;r++{
	heap.Push(&pq, &Item{r, nums[r]})
   }
   res := append([]int{}, pq[0].priority)
   for r<len(nums){
	l++
	for len(pq)>0 && pq[0].val < l {
		heap.Pop(&pq)
	}
	heap.Push(&pq, &Item{r, nums[r]})
   	res = append(res, pq[0].priority)
	r++
   }
   return res
}