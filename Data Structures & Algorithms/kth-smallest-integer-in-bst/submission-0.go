/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(*TreeNode)
	var res int
	dfs = func (root *TreeNode){
		if root == nil{
			return
		}
		dfs(root.Left)
		k--
		if k==0{
			res=root.Val
		}
		dfs(root.Right)	
	}
   dfs(root)
   return res
}

