package classical_150

// setZeros 矩阵置零
func setZeroes(matrix [][]int) {
	rowZeros := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				if i == 0 {
					rowZeros = 1
				} else {
					matrix[i][0] = 0
				}
			}
		}
	}
	for i := len(matrix) - 1; i > 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowZeros == 1 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}
