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

   
   slow.Next = nil
   curr := mid
   var prev *ListNode
   for curr != nil {
	   curr, curr.Next, prev = curr.Next, prev, curr 
   }


mid = prev
   

   // pop from stack and push to head 
   curr = head
   for mid!=nil {
      temp := mid.Next
	mid.Next = curr.Next
	curr.Next = mid
	curr = mid.Next
	mid =temp
   }
}
