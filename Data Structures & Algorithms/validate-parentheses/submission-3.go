func isValid(s string) bool {
   match := map[byte]byte{
	'[':']',
	'{':'}',
	'(':')',
   } 
   stack := make([]byte, 0, len(s))
   for i := range s{
	if _, ok := match[s[i]]; ok{
		stack = append(stack, s[i])
	} else {
		n := len(stack)
		if len(stack)==0 || match[stack[n-1]] != s[i]{
			return false
		}
		stack = stack[:len(stack)-1]
	}
   }
   return len(stack) == 0
}
