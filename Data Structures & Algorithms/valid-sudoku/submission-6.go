func isValidSudoku(board [][]byte) bool {
	m := make(map[byte]int)

	for rowID := 0; rowID < 9; rowID++{
		for i := 0; i < 9; i++{
			val := board[rowID][i]
			if val != '.' {
				m[val]++
			}
		}

		for _, val := range m {
			if val > 1 {
				return false
			}
		}
		clear(m)
	}
	clear(m)

	for columnID := 0; columnID < 9; columnID++ {
		for i := 0; i < 9; i++ {
			val := board[i][columnID]
			if val != '.'{
				m[val]++
			}
		}

		for _, val := range m{
			if val > 1 {
				return false
			}
		}
		clear(m)
	}
	clear(m)

	for box := 0; box < 9; box++{
		startRow := 3 * (box / 3)
		startColumn := 3 * (box % 3)

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++{
				val := board[startRow + i][startColumn + j]
				if val != '.'{
					m[val]++
				}
			}
		}

		for _, val := range m {
			if val > 1 {
				return false
			}
		}
		clear(m)
	}

	return true	
}



