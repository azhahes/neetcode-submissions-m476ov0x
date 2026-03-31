func lengthOfLongestSubstring(s string) int {
	dict := [128]int{}
	for i:=0; i<128; i++{
	dict[i] = -1
	}

	left, right:= 0,0
	res := 0
	for right < len(s){
		i := dict[s[right]]
		if i != -1 && i >= left {
			left = i + 1
		}
		dict[s[right]] = right
		res = max(res, right-left+1)
		right++
	}
	return res
}
