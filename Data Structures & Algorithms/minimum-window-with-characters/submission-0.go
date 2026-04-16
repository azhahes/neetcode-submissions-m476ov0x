// brute force
func minWindow(s string, t string) string {
	countT := make(map[byte]int)
	countW := make(map[byte]int)
	for _, c := range t{
		countT[byte(c)]++
	}
	l,r := 0,0
	resl, resr, resLen := 0, -1, math.MaxInt
	have, need := 0, len(countT)
	for r<len(s){
		countW[s[r]]++
		if countW[s[r]] == countT[s[r]]{
			have++
		}
		for have == need {
			if (r-l+1) < resLen{
				resl, resr, resLen = l, r, r-l+1
			}
			countW[s[l]]--
			if countW[s[l]] < countT[s[l]]{
				have--
			}
			l++
		}
		r++
	}
	return s[resl:resr+1]
}
