func isValidSudoku(board [][]byte) bool {
	for i:=0; i<9; i++{
		set := make(map[byte]bool)
		for j:=0; j<9; j++{
			if set[board[i][j]]{
				return false
			}
			if board[i][j] != '.'{
				set[board[i][j]] = true
			}
		}
	}

	for i:=0; i<9; i++{
		set := make(map[byte]bool)
		for j:=0; j<9; j++{
			if set[board[j][i]]{
				return false
			}
			if board[j][i] != '.'{
				set[board[j][i]] = true
			}
		}
	}

	for k:=0; k<9; k+=3{
		for l:=0; l<9; l+=3{
			set := make(map[byte]bool)
			for i:=k; i<k+3; i++{
				for j:=l; j<l+3; j++{
					if set[board[i][j]]{
						return false
					}
					if board[i][j] != '.'{
						set[board[i][j]] = true
					}
				}
			}
		}
	}
return true

}
