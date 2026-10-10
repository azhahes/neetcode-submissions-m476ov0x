/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }

3 -> 7 -> 4 -> 5 -> nil
^
3 -> n3 -> 7

 */

func copyRandomList(head *Node) *Node {
    curr := head
	for curr != nil {
		n := &Node{Val: curr.Val}
		n.Next = curr.Next
		curr.Next = n
		curr = n.Next
	}
	curr = head
	for curr != nil {
		if curr.Random != nil{
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}
	p := &Node{}
	newCurr := p
	curr = head
	for curr != nil {
		newCurr.Next = curr.Next
		curr.Next = curr.Next.Next
		curr = curr.Next
		newCurr = newCurr.Next
	}
	return p.Next
}
