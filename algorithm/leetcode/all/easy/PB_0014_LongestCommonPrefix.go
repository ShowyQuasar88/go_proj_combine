package easy

// longestCommonPrefix 最长公共前缀
func longestCommonPrefix(strs []string) string {
	idx := 0
	for ; idx < len(strs[0]); idx++ {
		same := true
		for j := 1; j < len(strs); j++ {
			if idx >= len(strs[j]) || strs[j][idx] != strs[0][idx] {
				same = false
				break
			}
		}
		if !same {
			return strs[0][:idx]
		}
	}
	return strs[0]
}
