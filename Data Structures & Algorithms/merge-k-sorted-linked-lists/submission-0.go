/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type PriorityQueue []*ListNode

func(pq PriorityQueue) Len() int { return len(pq)}
func(pq PriorityQueue) Less(i,j int) bool{ return pq[i].Val < pq[j].Val}
func(pq PriorityQueue) Swap(i,j int) { pq[i],pq[j]=pq[j],pq[i]}

func(pq *PriorityQueue) Push(v any) { *pq = append(*pq, v.(*ListNode))}

func(pq *PriorityQueue) Pop() any { 
	old := *pq
	n := len(old)
	val := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return val
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0, len(lists))

    for _, node := range lists {
        if node != nil {
            pq = append(pq, node)
        }
    }

   heap.Init(&pq)
   p := &ListNode{}
   curr := p 
   for len(pq) != 0 {
	v := heap.Pop(&pq).(*ListNode)
	curr.Next = v
	if v.Next != nil{
		heap.Push(&pq, v.Next)
	}
	curr = curr.Next
   }
   return p.Next

}
