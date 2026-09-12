func characterReplacement(s string, k int) int {
	var start, end, longest, maxCount int
	dict := make(map[byte]int)
	for end<len(s){
		dict[s[end]]++
		maxCount = max(maxCount, dict[s[end]])
		for start<len(s) && ((end-start+1)-maxCount) > k{
			dict[s[start]]--
			start++
		}
		longest = max(longest, end-start+1)
		end++
	}
	return longest
}
