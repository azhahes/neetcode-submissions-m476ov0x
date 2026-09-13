/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
   queue := []*TreeNode{root}
   res := make([][]int, 0)
   for len(queue) >0{
	n := len(queue)
	curr := make([]int, 0)
	for range n {
		top := queue[0]
		queue = queue[1:]
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
		curr = append(curr, top.Val)
	}
	res = append(res, curr)
   }
   return res
}
