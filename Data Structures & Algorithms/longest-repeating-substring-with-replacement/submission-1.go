func characterReplacement(s string, k int) int {
	maxW:=0
	for i:=0;i<len(s);i++{
		c := s[i]
		ct := 0
		for j:=i; j< len(s); j++{
			if c == s[j]{
				ct++
			}
			diff := (j-i+1)-ct 
			if diff <= k{
				maxW = max(maxW, j-i+1 + (k-diff))
			}
		}
	}
	return min(maxW, len(s))
}
