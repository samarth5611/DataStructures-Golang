const N = 9

func isValidSudoku(board [][]byte) bool {
	ans := true

	if hasInValidRow(board) == false {
		return false
	}
	if hasInValidColumn(board) == false {
		return false
	}
	if hasInValidSubMatrix(board) == false {
		return false
	}
	return ans
}

func hasInValidRow(a [][]byte) bool {
	for i := 0; i < N; i++ {
		mp := map[string]int{}
		for j := 0; j < N; j++ {
			if string(a[i][j]) == "." {
				continue
			}
			mp[string(a[i][j])]++
			if mp[string(a[i][j])] > 1 {
				return false
			}
		}
	}
	return true
}

func hasInValidColumn(a [][]byte) bool {
	for i := 0; i < N; i++ {
		mp := map[string]int{}
		for j := 0; j < N; j++ {
			if string(a[j][i]) == "." {
				continue
			}
			mp[string(a[j][i])]++
			if mp[string(a[j][i])] > 1 {
				return false
			}
		}
	}
	return true
}

func hasInValidSubMatrix(board [][]byte) bool {
	// map of maps in go
	mp := map[int]map[string]int{}
	for i := 0; i <= 8; i++ {
		mp[i] = make(map[string]int)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if string(board[i][j]) == "." {
				continue
			}
			groupIndex := getGroupIndex(i, j)
			mp[groupIndex][string(board[i][j])]++
			if mp[groupIndex][string(board[i][j])] > 1 {
				return false
			}
		}
	}
	return true
}

func getGroupIndex(x int, y int) int {
	val := 3*(x/3) + (y / 3)
	return val
}