package classical_150

func isSubsequence(s string, t string) bool {
	left, right := 0, 0
	for left < len(s) && right < len(t) {
		if s[left] == t[right] {
			left++
			right++
		} else {
			for right < len(t) && s[left] != t[right] {
				right++
			}
		}
	}
	if left == len(s) {
		return true
	}
	return false
}
