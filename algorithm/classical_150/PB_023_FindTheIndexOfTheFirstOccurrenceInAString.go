package classical_150

// strStr 找出字符串中第一个匹配项的下标
func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) || (len(haystack) == len(needle) && haystack != needle) {
		return -1
	}
	next, i, j := getNext(needle), 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if next[j] == -1 {
			j = 0
			i++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - len(needle)
	}
	return -1
}

func getNext(needle string) []int {
	next, b := make([]int, len(needle)), []byte(needle)
	next[0] = -1
	for curIdx, curMatch := 2, 0; curIdx < len(b); {
		if b[curIdx-1] == b[curMatch] {
			curMatch++
			next[curIdx] = curMatch
			curIdx++
		} else if curMatch > 0 {
			curMatch = next[curMatch]
		} else {
			curIdx++
		}
	}
	return next
}
