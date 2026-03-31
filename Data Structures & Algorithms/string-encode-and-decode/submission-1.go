type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, s := range strs {
		sb.WriteString(fmt.Sprintf("%d#", len(s)))
		sb.WriteString(s)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	cArr := []rune(encoded)
	res := make([]string, 0)
	i:=0
	for i<len(cArr) {
		var nums strings.Builder
		for cArr[i] != '#'{
			nums.WriteString(string(cArr[i]))
			i++
		}
		i++
		num, _ := strconv.Atoi(nums.String())
		
		var str strings.Builder
		for j := i; j<i+num; j++{
			str.WriteString(string(cArr[j]))
		}
		i += num
		res = append(res, str.String())
	}
	return res
}
