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
   l, r := 0, 0
   pq := make(PriorityQueue, 0)
   heap.Init(&pq)
   for r<k {
	heap.Push(&pq, &Item{r, nums[r]})
	r++
   }
   res := make([]int, 0)
   val := pq[0]
   res = append(res, val.priority)
   for r<len(nums){
	l++
	for len(pq) >0 {
   		val := pq[0]
		if val.val>=l && val.val <=r{
			break
		}
		heap.Remove(&pq, 0)
	}
	heap.Push(&pq, &Item{r, nums[r]})
   	val := pq[0].priority
   	res = append(res, val)
	r++
   }
   return res
}
/*
[1,3,2,1,-2,-3]

2,1,3
1,2,1,3

*/