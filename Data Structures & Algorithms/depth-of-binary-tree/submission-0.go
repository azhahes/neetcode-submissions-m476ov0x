/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
   maxDepth := 0
   var depthFunc func(*TreeNode, int)
   depthFunc = func(root *TreeNode, n int){
      if root == nil {
         return
      }
      depthFunc(root.Left, n+1)
      depthFunc(root.Right, n+1)
      maxDepth = max(maxDepth, n)
      n = n-1
   } 
   depthFunc(root, 1)
   return maxDepth
}
