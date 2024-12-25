package classical_150

// isHappy 快乐数
func isHappy(n int) bool {
	mp := make(map[int]struct{})
	for n != 1 {
		cur := 0
		for n != 0 {
			cur += (n % 10) * (n % 10)
			n /= 10
		}
		_, ok := mp[cur]
		if ok {
			return false
		}
		mp[cur] = struct{}{}
		n = cur
	}
	return true
}
