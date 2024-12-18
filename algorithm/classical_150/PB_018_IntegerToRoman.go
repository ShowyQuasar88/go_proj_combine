package classical_150

// intToRoman 整数转罗马
func intToRoman(num int) string {
	trans := [4][]string{
		{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"M", "MM", "MMM"},
	}
	res := ""
	if num/1000 > 0 {
		res += trans[3][num/1000-1]
		num %= 1000
	}
	if num/100 > 0 {
		res += trans[2][num/100-1]
		num %= 100
	}
	if num/10 > 0 {
		res += trans[1][num/10-1]
		num %= 10
	}
	if num > 0 {
		res += trans[0][num-1]
	}
	return res
}
