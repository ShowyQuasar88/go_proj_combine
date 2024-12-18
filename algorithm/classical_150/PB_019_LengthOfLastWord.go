package classical_150

// lengthOfLastWord 最后一个单词的长度
func lengthOfLastWord(s string) int {
	ans, idx := 0, len(s)-1
	for idx >= 0 && s[idx] == ' ' {
		idx--
	}
	if idx < 0 {
		return ans
	}
	for idx >= 0 && s[idx] != ' ' {
		ans++
		idx--
	}
	return ans
}
