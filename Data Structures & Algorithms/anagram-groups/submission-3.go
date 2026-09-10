import (
	"slices"
)
func groupAnagrams(strs []string) [][]string {
	dict :=  make(map[string][]string)
	for _, str := range strs {
		arr := []byte(str)
		slices.Sort(arr)
		dict[string(arr)] = append(dict[string(arr)] ,str)  
	}
	result := make([][]string,0, len(dict))
	for _, val := range dict {
		result = append(result, val)
	}
	return result
}
