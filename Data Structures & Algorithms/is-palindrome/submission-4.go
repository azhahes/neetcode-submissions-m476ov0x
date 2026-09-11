func IsAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start<end{
		for start<end && !IsAlphaNumeric(rune(s[start])){start++}
		for start<end && !IsAlphaNumeric(rune(s[end])){end--}
		if strings.ToUpper(string(s[start])) != strings.ToUpper(string(s[end])){
			return false
		}
		start++
		end--
	}
	return true
}
