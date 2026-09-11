type Item struct{
	val int
	count int
}

type PriorityQueue []*Item

func(pq PriorityQueue) Len() int {
	return len(pq)
}

func(pq PriorityQueue) Less(i,j int) bool{
	return pq[i].count>pq[j].count
}

func(pq PriorityQueue) Swap(i,j int) {
	pq[i],pq[j] = pq[j], pq[i]
}

func(pq *PriorityQueue) Push(i any) {
	item := i.(*Item)
	*pq = append(*pq, item)
}
func(pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item

}


func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n]++
	}
	pq := make(PriorityQueue,0, len(countMap))
	for k,v := range countMap {
		pq = append(pq, &Item{k,v})
	}	
	heap.Init(&pq)
	result := make([]int,0,k)
	for range k {
		i := heap.Pop(&pq).(*Item)
		result = append(result, i.val)
	}
	return result
}
