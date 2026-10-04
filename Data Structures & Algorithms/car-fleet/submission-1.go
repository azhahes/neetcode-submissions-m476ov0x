func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int, len(position))
	for i:=0;i<len(position);i++{
		cars[i] = [2]int{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0]>cars[j][0]
	})
	stack := make([]float64, 0)
	for _, v := range cars{
		time := float64((target-v[0]))/float64(v[1])
		if len(stack) == 0 || stack[len(stack)-1] < time {
			stack = append(stack, time)
		}
	}
	return len(stack)
}
