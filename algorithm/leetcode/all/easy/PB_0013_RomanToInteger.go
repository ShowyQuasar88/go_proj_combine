package easy

// romanToInt 罗马数字转整数
func romanToInt(s string) int {
	help, b := make([]int, len(s)), []byte(s)
	for idx, each := range b {
		switch each {
		case 'I':
			help[idx] = 1
		case 'V':
			help[idx] = 5
		case 'X':
			help[idx] = 10
		case 'L':
			help[idx] = 50
		case 'C':
			help[idx] = 100
		case 'D':
			help[idx] = 500
		case 'M':
			help[idx] = 1000
		default:
			help[idx] = 0
		}
	}
	ans := help[len(help)-1]
	for i := len(help) - 2; i >= 0; i-- {
		if help[i] < help[i+1] {
			ans -= help[i]
		} else {
			ans += help[i]
		}
	}
	return ans
}
