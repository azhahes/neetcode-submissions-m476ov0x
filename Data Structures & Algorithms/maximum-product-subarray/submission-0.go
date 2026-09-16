import(
	"slices"
)
func maxProduct(nums []int) int {
   m, n := 1,1
   res := slices.Max(nums)
   for _, num := range nums {
	m, n = max(num*m, num*n, num), min(num*m, num*n, num)
	res = max(res, m)
   }
   return res
}
