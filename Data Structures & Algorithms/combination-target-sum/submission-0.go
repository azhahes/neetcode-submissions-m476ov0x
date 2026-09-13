import(
	"slices"
)

func combinationSum(nums []int, target int) [][]int {
   result := make([][]int, 0)
   var dfs func(int, int, []int)
	dfs = func (i,total int, curr []int) {
		if total == target{
			result = append(result, slices.Clone(curr))
			return
		}
	
		if i>= len(nums) || total>target{
			return
		}

		curr = append(curr, nums[i])
		dfs(i,total+nums[i], curr )
		curr = curr[:len(curr)-1]
		dfs(i+1,total, curr )
	}
	dfs(0,0,[]int{})
	return result
}


