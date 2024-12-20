package classical_150

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			count := cellAlive(board, i, j)
			if board[i][j] == 1 && (count < 2 || count > 3) {
				board[i][j] = -1
			} else if board[i][j] == 0 && count == 3 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func cellAlive(board [][]int, x, y int) int {
	count := 0
	if x > 0 && y > 0 && (board[x-1][y-1] == 1 || board[x-1][y-1] == -1) {
		count++
	}
	if x > 0 && (board[x-1][y] == 1 || board[x-1][y] == -1) {
		count++
	}
	if x > 0 && y < len(board[0])-1 && (board[x-1][y+1] == 1 || board[x-1][y+1] == -1) {
		count++
	}
	if y > 0 && (board[x][y-1] == 1 || board[x][y-1] == -1) {
		count++
	}
	if y < len(board[0])-1 && (board[x][y+1] == 1 || board[x][y+1] == -1) {
		count++
	}
	if y > 0 && x < len(board)-1 && (board[x+1][y-1] == 1 || board[x+1][y-1] == -1) {
		count++
	}
	if x < len(board)-1 && (board[x+1][y] == 1 || board[x+1][y] == -1) {
		count++
	}
	if x < len(board)-1 && y < len(board[0])-1 && (board[x+1][y+1] == 1 || board[x+1][y+1] == -1) {
		count++
	}
	return count
}
