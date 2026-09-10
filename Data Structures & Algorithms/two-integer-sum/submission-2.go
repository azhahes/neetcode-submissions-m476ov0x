func twoSum(nums []int, target int) []int {
   dict := make(map[int]int)
   for i, num := range nums {
	if val, ok := dict[target - num]; ok {
		return []int{val, i}
	} 
	dict[num] = i
   }
   return nil
}
