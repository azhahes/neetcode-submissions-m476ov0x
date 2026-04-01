/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
   var prev *ListNode
   curr := head
   for curr != nil{
	 curr, curr.Next, prev = curr.Next, prev, curr
   } 
   return prev
}
