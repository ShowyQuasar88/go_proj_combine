package easy

import "strconv"

// isPalindrome 回文数
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for left, right := 0, len(s)-1; left <= right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}
