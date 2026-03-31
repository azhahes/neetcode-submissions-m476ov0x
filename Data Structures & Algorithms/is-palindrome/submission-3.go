func isValid(r byte) bool {
	return (r >= 'a' && r <= 'z') || 
	(r >= 'A' && r <= 'Z') || 
	(r >= '0' && r <= '9')
}
func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		for start < end && !isValid(s[start]) {
			start++
		}
		for start < end && !isValid(s[end]) {
			end--
		}


		if strings.ToUpper(string(s[start])) != strings.ToUpper(string(s[end])) {
			return false
		}
		start++
		end--
	}
	return true

}