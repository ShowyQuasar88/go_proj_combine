package classical_150

// Candy 分发糖果
func Candy(ratings []int) int {
	tmp, length := make([]int, len(ratings)), len(ratings)
	if length == 1 {
		return 1
	}
	for i := 0; i < length; i++ {
		tmp[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && tmp[i] <= tmp[i-1] {
			tmp[i] = tmp[i-1] + 1
		}
	}
	res := tmp[length-1]
	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && tmp[i] <= tmp[i+1] {
			tmp[i] = tmp[i+1] + 1
		}
		res += tmp[i]
	}
	return res
}
