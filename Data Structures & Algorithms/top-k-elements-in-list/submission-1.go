type Item struct{
	num int
	count int
}

type PriorityQueue []*Item

func(p PriorityQueue) Len() int {
	return len(p)
}

func(p PriorityQueue) Less(i,j int) bool {
	return p[i].count > p[j].count
}

func(p PriorityQueue) Swap(i,j int) {
	p[i], p[j] = p[j], p[i]
}

func(p *PriorityQueue) Push(v any) {
	val := v.(*Item)
	*p = append(*p, val)
}

func(p *PriorityQueue) Pop() any {
	n := len(*p)
	v := (*p)[n-1]
	*p = (*p)[:n-1]
	return v
}



func topKFrequent(nums []int, k int) []int {
	countM := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := countM[n]; ok{
			countM[n]++
		} else {
			countM[n] = 1
		}
	}
	pq := make(PriorityQueue, 0, len(nums))
	heap.Init(&pq)
	for k, v := range countM {
		item := &Item{k, v}
		heap.Push(&pq, item)
	}
	res:= make([]int,0, k)
	for i := 0; i < k; i++{
		v := heap.Pop(&pq).(*Item)
		res = append(res, v.num)
	}
	return res

}
