/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	p := &ListNode{}
	curr := p 
	for l1 != nil || l2 != nil {
		n := carry
		if l1!= nil{
			n += l1.Val
			l1 = l1.Next
		}
		if l2!= nil{
			n += l2.Val
			l2 = l2.Next
		}
		sum := n%10
		carry = n/10
		curr.Next = &ListNode{Val:sum}
		curr = curr.Next
	}
	if carry > 0{
		curr.Next = &ListNode{Val:carry}
	}
	return p.Next

}
