func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, r := range s{
		sMap[r]++
	}
	for _, r := range t{
		tMap[r]++
	}

	if len(sMap) != len(tMap){
		return false
	}

	for k, v := range sMap {
		if tMap[k] != v{
			return false
		}
	}
	return true

}
