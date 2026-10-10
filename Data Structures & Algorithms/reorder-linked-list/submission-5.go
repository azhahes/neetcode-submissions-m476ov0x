/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
curr = 8 -> 6 -> nil
head = 2 -> 4 -> 6 -> 8 -> nil

 */

func reorderList(head *ListNode) {
	mid := getMidNode(head)
	curr := reverseList(mid)
	for curr != nil {
		temp := curr.Next
		curr.Next = head.Next
		head.Next = curr
		head = curr.Next
		curr = temp
	}
	head.Next = nil
}

func reverseList(head *ListNode)*ListNode{
	var prev *ListNode
	curr := head
	for curr != nil {
		curr, curr.Next, prev = curr.Next, prev, curr
	}
	return prev
}

func getMidNode(head *ListNode)*ListNode{
	p := &ListNode{Val:0, Next:head}
	slow, fast := head, p 
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
