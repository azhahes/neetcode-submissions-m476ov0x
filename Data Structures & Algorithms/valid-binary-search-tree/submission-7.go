/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, low int, high int) bool {
		if root == nil {
			return true
		}
		if root.Val <= low || root.Val >= high{
			return false
		}
		return dfs(root.Left, low, root.Val) && dfs(root.Right, root.Val, high)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
