package classical_150

// romanToInt 罗马数字转换
func romanToInt(s string) int {
	length := len(s)
	ans := make([]int, length)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			ans[i] = 1
		case 'V':
			ans[i] = 5
		case 'X':
			ans[i] = 10
		case 'L':
			ans[i] = 50
		case 'C':
			ans[i] = 100
		case 'D':
			ans[i] = 500
		case 'M':
			ans[i] = 1000
		}
	}
	res := ans[length-1]
	for i := length - 2; i >= 0; i-- {
		if ans[i] < ans[i+1] {
			res -= ans[i]
		} else {
			res += ans[i]
		}
	}
	return res
}
