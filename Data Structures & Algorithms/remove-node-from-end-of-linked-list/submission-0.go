/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
   stack := make([]*ListNode, 0) 
   curr := head
   for curr != nil {
      stack = append(stack, curr)
      curr = curr.Next
   }
   
   length := len(stack)
    // Edge case: remove the head node
    if n == length {
        return head.Next
    }

   curr = stack[length-n-1]
   curr.Next = curr.Next.Next
   return head
}
