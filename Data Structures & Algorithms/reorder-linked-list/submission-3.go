/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next ==nil {
		return
	}
   slow, fast := head, head.Next
   for fast != nil && fast.Next != nil {
	slow = slow.Next
	fast = fast.Next.Next
   } 
   mid := slow.Next
   
   stack := make([]*ListNode,0)
   for mid != nil {
	stack = append(stack, mid)
	mid = mid.Next
   }

   curr := head
   for len(stack)>0 {
	n := len(stack)
	top := stack[n-1]
	stack = stack[:n-1]
	curr.Next, top.Next = top, curr.Next
	curr = top.Next
   }
   curr.Next = nil
}
