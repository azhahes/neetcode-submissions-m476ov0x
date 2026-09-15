func longestPalindrome(s string) string {
	resL, resR, resLen := 0,-1,0
	for i := range s{
		l, r := i, i
		for l>=0 && r<len(s) && s[l] == s[r] {
			if (r-l+1)>resLen{
				resL, resR, resLen = l, r, (r-l+1)
			}
			l--
			r++
		}

		l, r = i, i+1
		for l>=0 && r<len(s) && s[l] == s[r] {
			if (r-l+1)>resLen{
				resL, resR, resLen = l, r, (r-l+1)
			}
			l--
			r++
		}
	}
	return s[resL:resR+1]
}
