/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var maxD int
func diameterOfBinaryTree(root *TreeNode) int {
   maxD=0
   dfs(root)
   return maxD
}
func dfs(root *TreeNode) int {
   if root == nil {
	   return 0
   } 
   left := dfs(root.Left)
   right := dfs(root.Right)
   maxD = max(maxD, left+right) 
   return 1 + max(left, right)
}
