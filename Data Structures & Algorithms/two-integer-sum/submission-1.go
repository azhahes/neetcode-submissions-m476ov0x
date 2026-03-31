func twoSum(nums []int, target int) []int {
   dict := map[int]int{}
   for i, n := range nums {
	if v, ok := dict[target-n];ok{
		return []int{v,i}
	}
	dict[n] = i
   } 
   return nil
}
