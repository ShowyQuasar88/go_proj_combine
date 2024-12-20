package classical_150

// rotateImage 旋转图像
func rotateImage(matrix [][]int) {
	for top, bottom := 0, len(matrix)-1; top < bottom; top, bottom = top+1, bottom-1 {
		for begin := 0; top+begin < bottom; begin++ {
			tmp := matrix[top][top+begin]
			matrix[top][top+begin] = matrix[bottom-begin][top]
			matrix[bottom-begin][top] = matrix[bottom][bottom-begin]
			matrix[bottom][bottom-begin] = matrix[top+begin][bottom]
			matrix[top+begin][bottom] = tmp
		}
	}
}
