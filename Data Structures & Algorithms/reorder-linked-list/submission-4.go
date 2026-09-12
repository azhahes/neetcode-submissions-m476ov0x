/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
   // find middle
   slow, fast := head, head.Next
   for fast != nil && fast.Next != nil {
	slow = slow.Next
	fast = fast.Next.Next
   }
   mid := slow.Next
   // from middle push to stack
   stack := make([]*ListNode, 0)
   for mid != nil {
	stack = append(stack, mid)
	mid = mid.Next
   }
   // pop from stack and push to head 
   curr := head
   for len(stack)>0 {
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	top.Next = curr.Next
	curr.Next = top
	curr = top.Next
   }
   curr.Next = nil
}
