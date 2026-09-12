func minWindow(s string, t string) string {
   countT := make(map[byte]int) 
   countW := make(map[byte]int) 
   for i := range t {
	countT[t[i]]++
   }
   var resl, resr, resLen int
   var l, r, have, need int
   need = len(countT)
   resLen = math.MaxInt
   resr = -1
   for r<len(s){
	countW[s[r]]++
	if countW[s[r]] == countT[s[r]]{
		have++
	}
	for have == need{
		if (r-l+1) < resLen{
			resl, resr, resLen = l, r, (r-l+1)
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
