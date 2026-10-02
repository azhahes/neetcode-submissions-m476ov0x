func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start<end{
		currSum := numbers[start] + numbers[end]
		if target == currSum{
			return []int{start+1, end+1}
		}
		if currSum>target{
			end--
		}else {
			start++
		}
	}
	return nil
}
