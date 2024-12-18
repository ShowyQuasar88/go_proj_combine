package classical_150

// longestCommonPrefix 最长公共前缀
func longestCommonPrefix(strs []string) string {
	length, ans := len(strs), ""
	if length == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs[0]); i++ {
		flag := true
		for j := 1; j < length; j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		ans += string(strs[0][i])
	}
	return ans
}
