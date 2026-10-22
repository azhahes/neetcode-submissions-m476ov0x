/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
   return Codec{} 
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
   var dfs func(*TreeNode)
   res := make([]string,0)
   dfs = func(root *TreeNode){
	if root == nil {
		res = append(res, "N")
		return
	}
	res = append(res, strconv.Itoa(root.Val))
	dfs(root.Left)
	dfs(root.Right)
   }
   dfs(root)
   return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	i := 0
	var dfs func()*TreeNode
	dfs = func() *TreeNode{
		if arr[i] == "N" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(arr[i])
		root := &TreeNode{Val:val}
		i++
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
   return dfs()
}
