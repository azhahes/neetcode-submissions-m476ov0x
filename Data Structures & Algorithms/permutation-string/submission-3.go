func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1){
		return false
	}
	s1Count := [26]int{}
	s2Count := [26]int{}
	for i:=0;i<len(s1);i++{
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}
	matches := 0
	for i:=0;i<26;i++{
		if s1Count[i] == s2Count[i]{
			matches++
		}
	}
	l:=0
	for r:=len(s1);r<len(s2);r++{
		if matches == 26{
			return true
		}
		s2Count[s2[r]-'a']++
		if s1Count[s2[r]-'a'] == s2Count[s2[r]-'a']{
			matches++
		} else if s1Count[s2[r]-'a']+1 == s2Count[s2[r]-'a']{
			matches--
		}
		s2Count[s2[l]-'a']--
		if s1Count[s2[l]-'a'] == s2Count[s2[l]-'a']{
			matches++
		} else if s1Count[s2[l]-'a']-1 == s2Count[s2[l]-'a']{
			matches--
		}
		l++
	}
	return matches == 26
}
