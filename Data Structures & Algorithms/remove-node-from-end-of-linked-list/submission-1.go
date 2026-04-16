/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
   psudoHead := &ListNode{0,head}
    slow, fast := psudoHead, head
    for range n {
      fast = fast.Next
    }
    for fast != nil {
      slow = slow.Next
      fast = fast.Next
    }
    slow.Next = slow.Next.Next
    return psudoHead.Next
}
