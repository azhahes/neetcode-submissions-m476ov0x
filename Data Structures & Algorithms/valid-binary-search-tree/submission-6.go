/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type Node struct {
	node *TreeNode
	low int
	high int	
 }

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
   q := make([]*Node, 0)
	q = append(q, &Node{root, math.MinInt, math.MaxInt})
   for len(q) >0 {
	n := len(q)
	for i:=0;i<n;i++{
		top := q[0]
		q = q[1:]
		if top.node.Val <= top.low || top.node.Val >= top.high{
			return false
		}
		if top.node.Left != nil {
			q = append(q, &Node{top.node.Left, top.low, top.node.Val} )
		}
		if top.node.Right != nil {
			q = append(q, &Node{top.node.Right, top.node.Val, top.high})
		}
	}
   }
   return true
}
