func numIslands(grid [][]byte) int {
	r, c := len(grid), len(grid[0])
	res := 0
    v := make([][]bool, r)
	for i := range v {
		v[i] = make([]bool, c)
	} 

	q := make([][2]int,0)

	for m:= 0; m<r; m++{
		for n:=0; n<c; n++{
			if v[m][n] || grid[m][n] == '0' {
				continue	
			}
			v[m][n] = true
			q = append(q, [2]int{m,n})
			for len(q) > 0 {
				i, j := q[0][0], q[0][1]
				q = q[1:]
				if i+1 < r && grid[i+1][j] == '1' && !v[i+1][j] {
					v[i+1][j] = true
					q = append(q, [2]int{i+1,j})
				}
				if j+1 < c && grid[i][j+1] == '1' && !v[i][j+1] {
					v[i][j+1] = true
					q = append(q, [2]int{i,j+1})
				}
				if i-1 >= 0 && grid[i-1][j] == '1' && !v[i-1][j] {
					v[i-1][j] = true
					q = append(q, [2]int{i-1,j})
				}
				if j-1 >= 0 && grid[i][j-1] == '1' && !v[i][j-1] {
					v[i][j-1] = true
					q = append(q, [2]int{i,j-1})
				}
			}
			res++
		}
	}
	return res
}
