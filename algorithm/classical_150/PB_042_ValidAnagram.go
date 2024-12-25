package classical_150

// isAnagram 有效的字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	c, tmp := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		tmp[s[i]-'a']++
		tmp[t[i]-'a']--
	}
	return c == tmp
}
