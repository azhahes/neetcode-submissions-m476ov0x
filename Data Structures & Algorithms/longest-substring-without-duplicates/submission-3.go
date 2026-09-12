func lengthOfLongestSubstring(s string) int {
	dict := make(map[byte]int)
	var start, end, longest int
	for end < len(s){
		if v, ok := dict[s[end]]; ok{
			start = max(start, v+1)
		}
		longest = max(longest, end-start+1)
		dict[s[end]] = end
		end++
	}
	return longest
}
