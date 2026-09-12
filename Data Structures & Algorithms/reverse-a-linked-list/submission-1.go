/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
   curr:= head
   var prev *ListNode
   for curr != nil {
	curr, curr.Next, prev = curr.Next, prev, curr 
   }
   return prev
}
