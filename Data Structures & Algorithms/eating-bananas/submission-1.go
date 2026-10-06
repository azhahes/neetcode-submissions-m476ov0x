import (
	"slices"
)
func minEatingSpeed(piles []int, h int) int {
	var feasible func(int) bool 
	feasible = func(i int) bool {
		curr := 0
		for _, p := range piles {
			curr += ((p-1)/i)+1
			if curr > h{
				return false
			}
		}
		return true
	}
	l, r := 1, slices.Max(piles)
	for l<r {
		mid := l + (r-l)/2
		if feasible(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
