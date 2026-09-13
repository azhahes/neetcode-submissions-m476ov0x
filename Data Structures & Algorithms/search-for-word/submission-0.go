func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	set := make(map[[2]int]bool)
	var dfs func(int, int, int) bool
	dfs = func(r,c,i int) bool {
		if i == len(word){
			return true
		}
		if r<0 || c<0 || r>=n || c >=m || board[r][c] != word[i] || set[[2]int{r,c}]{
			return false
		}
		set[[2]int{r,c}] = true
		res := dfs(r+1,c,i+1) ||
			   dfs(r-1,c,i+1) ||	
			   dfs(r,c+1,i+1) ||	
			   dfs(r,c-1,i+1) 
		set[[2]int{r,c}] = false
		return res
	} 
	for i := range board {
		for j := range board[0]{
			if dfs(i,j,0){
				return true
			}
		}
	}
	return false
}
