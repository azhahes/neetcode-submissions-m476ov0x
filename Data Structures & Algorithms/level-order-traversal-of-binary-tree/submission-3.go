/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
   res := make([][]int, 0)
   if root == nil {
	return res
   }
   q := []*TreeNode{root}
   for len(q) > 0{
	curr := make([]int, 0)
	i := len(q)
	for i > 0{
		top := q[0]
		q = q[1:]
		curr = append(curr, top.Val)
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
		i--
	}
	res = append(res, curr)
   } 
   return res
}
