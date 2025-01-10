package easy

// isValid 有效的括号
func isValid(s string) bool {
	stack, idx := make([]byte, len(s)), -1
	for _, each := range []byte(s) {
		if each == '(' || each == '[' || each == '{' {
			idx++
			stack[idx] = each
		} else {
			if idx == -1 {
				return false
			}
			switch each {
			case ')':
				if stack[idx] != '(' {
					return false
				}
			case ']':
				if stack[idx] != '[' {
					return false
				}
			case '}':
				if stack[idx] != '{' {
					return false
				}
			}
			idx--
		}
	}
	return idx == -1
}
