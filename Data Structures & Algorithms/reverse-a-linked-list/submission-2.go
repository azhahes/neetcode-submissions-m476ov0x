/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
   prev = nil
   curr := head
   for curr != nil {
	curr, prev, curr.Next = curr.Next, curr, prev
   }
   return prev
}
