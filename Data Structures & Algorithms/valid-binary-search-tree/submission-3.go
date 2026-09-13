/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate (r *TreeNode, min, max int) bool{
	if r == nil {
		return true 
	}
	if r.Val >= max || r.Val <= min{
		return false
	}
	return validate(r.Left, min, r.Val) && validate(r.Right, r.Val, max)
}