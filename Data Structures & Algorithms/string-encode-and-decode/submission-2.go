type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result strings.Builder
	for _, s := range strs {
		result.WriteString(fmt.Sprintf("%d#%s", len(s), s))
	}
	return result.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string, 0)
	i := 0
	n := ""
	for i<len(encoded) {
		if encoded[i] == '#'{
			l, _ := strconv.Atoi(n)	
			result = append(result, encoded[i+1:i+l+1])
			i = i+l+1
			n = ""
		} else {
			n += string(encoded[i])
			i++
		}
	}
	return result
}
