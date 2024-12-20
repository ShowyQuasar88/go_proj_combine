package classical_150

// isIsomorphic 同构字符串
func isIsomorphic(s string, t string) bool {
	trans, diTrans := make(map[byte]byte), make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if tran, ok := trans[s[i]]; ok {
			if tran != t[i] {
				return false
			}
		} else {
			if diTrans[t[i]] {
				return false
			}
			diTrans[t[i]] = true
			trans[s[i]] = t[i]
		}
	}
	return true
}
