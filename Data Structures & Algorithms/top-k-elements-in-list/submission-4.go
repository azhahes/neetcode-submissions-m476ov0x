type PriorityQueue [][2]int

func(p PriorityQueue) Len() int { return len(p)}
func(p PriorityQueue) Less(i, j int) bool{ return p[i][0]>p[j][0]}
func(p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i]}

func(p *PriorityQueue) Push(a any) {
	*p = append(*p, a.([2]int))
}

func(p *PriorityQueue) Pop() any {
	val := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return val
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n]++
	}
	pq := make(PriorityQueue, len(countMap))
	i := 0
	for k,v := range countMap{
		pq[i] = [2]int{v,k}
		i++
	}
	heap.Init(&pq)
	res := make([]int, k)
	for i:=0;i<k;i++{
		res[i] = heap.Pop(&pq).([2]int)[1]
	}
	return res
}
