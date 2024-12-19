package classical_150

// spiralOrder 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	leftRow, leftCol, rightRow, rightCol := 0, 0, len(matrix)-1, len(matrix[0])-1
	for leftRow <= rightRow && leftCol <= rightCol {
		if leftRow == rightRow {
			for leftCol <= rightCol {
				res = append(res, matrix[leftRow][leftCol])
				leftCol++
			}
		} else if leftCol == rightCol {
			for leftRow <= rightRow {
				res = append(res, matrix[leftRow][leftCol])
				leftRow++
			}
		} else {
			res = spiral(matrix, leftRow, leftCol, rightRow, rightCol, res)
			leftRow++
			leftCol++
			rightRow--
			rightCol--
		}
	}
	return res
}

func spiral(matrix [][]int, leftRow, leftCol, rightRow, rightCol int, res []int) []int {
	initLeftCol, initLeftRow := leftCol, leftRow
	for leftCol < rightCol {
		res = append(res, matrix[leftRow][leftCol])
		leftCol++
	}
	for leftRow < rightRow {
		res = append(res, matrix[leftRow][rightCol])
		leftRow++
	}
	for rightCol > initLeftCol {
		res = append(res, matrix[rightRow][rightCol])
		rightCol--
	}
	for rightRow > initLeftRow {
		res = append(res, matrix[rightRow][initLeftCol])
		rightRow--
	}
	return res
}
