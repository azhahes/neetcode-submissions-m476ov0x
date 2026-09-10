func isAnagram(s string, t string) bool {
	var counter [26]int
	for _, r := range s{
		counter[int(r) - int('a')]++
	}
	for _, r := range t{
		counter[int(r) - int('a')]--
	}
	var empty [26]int
	return counter == empty 
}
