func groupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}
	for i, str := range strs {
		r := []rune(str)
		sort.Slice(r, func(i,j int) bool {
			return r[i]<r[j]
		})
		s := string(r)
		dict[s] = append(dict[s], strs[i])
	}
	result := [][]string{}
	for _,v := range dict {
		result = append(result, v)
	}
	return result
}
