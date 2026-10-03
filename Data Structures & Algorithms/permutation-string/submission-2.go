func checkInclusion(s1 string, s2 string) bool {
	if len(s2)< len(s1){
		return false
	}
	dict1 := make(map[byte]int)
	dict2 := make(map[byte]int)
	for i := range s1 {
		dict1[s1[i]]++
	}

	l, r := 0, 0
	for r< len(s2){
		if _, ok := dict1[s2[r]]; !ok{
			dict2 = make(map[byte]int)
			l=r+1
		} else {
			dict2[s2[r]]++
			for dict2[s2[r]] > dict1[s2[r]]{
				dict2[s2[l]]--
				l++
			} 
			if len(s1) == (r-l+1) {
				return true
			}
		}
		r++
	}
	return false
}
