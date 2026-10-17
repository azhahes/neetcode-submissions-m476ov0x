/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	var dfs func(*TreeNode, int)
	res := make([]int,0)
	dfs = func(root *TreeNode, depth int){
		if root == nil {
			return
		}
		if len(res) == depth{
			res = append(res, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 0)
	return res
}
