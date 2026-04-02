func characterReplacement(s string, k int) int {
	var l, r, res, maxf int
	count := make(map[byte]int)
	for r< len(s) {
		count[s[r]]++
		maxf = max(maxf, count[s[r]])
		if (r-l+1) - maxf > k {
			count[s[l]]--
			l++
		} 
		res = max(res, r-l+1)
		r++
	}
	return res
}
