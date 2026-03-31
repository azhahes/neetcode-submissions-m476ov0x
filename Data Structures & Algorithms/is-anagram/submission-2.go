// sorting
func isAnagram(s string, t string) bool {
	sArr := []byte(s)
	tArr := []byte(t)
	sort.Slice(sArr, func(i, j int) bool { return sArr[i]<sArr[j]})
	sort.Slice(tArr, func(i, j int) bool { return tArr[i]<tArr[j]})
	return string(sArr) == string(tArr)
}
