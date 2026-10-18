/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var dfs func(*TreeNode,int)
	res := 0
	dfs = func(root *TreeNode, prevMax int){
		if root == nil {
			return
		}
		if root.Val >= prevMax {
			res++
		}
		prevMax = max(prevMax, root.Val)
		dfs(root.Left, prevMax)
		dfs(root.Right, prevMax)
	}
	dfs(root,math.MinInt)
	return res
}
