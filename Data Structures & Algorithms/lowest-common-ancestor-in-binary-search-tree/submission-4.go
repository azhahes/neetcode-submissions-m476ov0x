/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

		   6
		 2    8
	   0  4  7  9
	     3 5
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
   curr := root
   if p.Val > q.Val {
	p, q = q, p
   } 
   for curr != nil {
	if curr.Val >= p.Val && curr.Val <= q.Val{
		return curr
	}
	if curr.Val > q.Val {
		curr = curr.Left
	} else {
		curr = curr.Right
	}
   }
   return curr
}
