package classical_150

// canConstruct 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	counts := [26]int{}
	for _, r := range magazine {
		counts[r-'a']++
	}
	for _, r := range ransomNote {
		if counts[r-'a'] == 0 {
			return false
		}
		counts[r-'a']--
	}
	return true
}
