/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
   dict := make(map[*Node]*Node)
   curr := head
   for curr != nil {
	newNode := &Node{Val:curr.Val}
	dict[curr] = newNode
	curr = curr.Next
   } 
   curr = head
   for curr != nil {
	n := dict[curr]
	nxt := dict[curr.Next]
	r := dict[curr.Random]
	n.Next = nxt
	n.Random = r
	curr = curr.Next
   }
   return dict[head]
}
