/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
   res := true
   var dfs func(*TreeNode) int
   dfs = func(node *TreeNode) int{
	if node == nil {
		return 0
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	if abs(left-right) > 1{
		res = false
	}
	return 1 + max(left, right)
   }
   dfs(root)
   return res
}
func abs(v int) int{
	if v<0{
		return v*-1
	}
	return v
}
