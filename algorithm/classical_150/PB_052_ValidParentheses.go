package classical_150

// isValid 有效的括号
func isValid(s string) bool {
	stack, top, sb := make([]byte, len(s)), -1, []byte(s)
	for _, b := range sb {
		switch b {
		case '[', '(', '{':
			top++
			stack[top] = b
		case ']':
			if top == -1 || stack[top] != '[' {
				return false
			}
			top--
		case ')':
			if top == -1 || stack[top] != '(' {
				return false
			}
			top--
		case '}':
			if top == -1 || stack[top] != '{' {
				return false
			}
			top--
		}
	}
	return top == -1
}
