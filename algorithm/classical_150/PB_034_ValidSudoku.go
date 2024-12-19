package classical_150

import "strconv"

// isValidSudoku 有效的数独
func isValidSudoku(board [][]byte) bool {
	rows, cols, blocks := make([][9]bool, 9), make([][9]bool, 9), make([][9]bool, 9)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			char := board[row][col]
			if char == '.' {
				continue
			}
			cur, _ := strconv.Atoi(string(char))
			block := (row/3)*3 + col/3
			if rows[row][cur-1] || cols[col][cur-1] || blocks[block][cur-1] {
				return false
			}
			rows[row][cur-1] = true
			cols[col][cur-1] = true
			blocks[block][cur-1] = true
		}
	}
	return true
}
