func findDuplicate(nums []int) int {
    slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow2 := 0 
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2{
			return slow
		}
	}
	return -1
}

/*
0 -> 1 -> 2 -> 3 -> 2 
		  ^	        |
		  |<- <- <- V

2d = 2(s+c)
*/