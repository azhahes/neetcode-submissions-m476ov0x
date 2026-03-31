func isAnagram(s string, t string) bool {
	dict1 := [26]int{}
	dict2 := [26]int{}
	for _, c := range s{
		dict1[c - 'a']++
	}

	for _, c := range t{
		dict2[c - 'a']++
	}

	for i := range 26 {
		if dict1[i] != dict2[i]{
			return false
		}
	}
	return true
}
