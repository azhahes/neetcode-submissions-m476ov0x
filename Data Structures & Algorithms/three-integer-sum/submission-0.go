import("maps")
func threeSum(nums []int) [][]int {
	resultSet := make(map[[3]int]bool)
	for i, n := range nums{
		target := 0-n
		dict := make(map[int]bool)
		for j:= i+1; j<len(nums); j++{
			if _, ok:= dict[target-nums[j]]; ok{
				val := []int{nums[j],n, target-nums[j]}
				sort.Ints(val)
				resultSet[[3]int(val)] = true
			}
			dict[nums[j]] = true
		}
	}
	result := make([][]int,0)
	for  v:=range maps.Keys(resultSet){
		result = append(result, v[:])
	}
	return result
}