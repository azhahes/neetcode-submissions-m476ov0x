/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
   if p.Val > q.Val{
	q,p = p,q
   }
   curr := root
   for curr != nil {
	if curr.Val >= p.Val && curr.Val <= q.Val{
		return curr
	}
	if curr.Val > q.Val{
		curr = curr.Left
	} else{
		curr = curr.Right
	}
   }
   return root
}
