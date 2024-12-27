package classical_150

import "strings"

// simplifyPath 简化路径
func simplifyPath(path string) string {
	stack, top, ans := make([]string, len(path)), -1, ""
	for _, p := range strings.Split(path, "/") {
		if p == ".." {
			if top != -1 {
				top--
			}
		} else if p != "" && p != "." {
			top++
			stack[top] = p
		}
	}
	for i := 0; i <= top; i++ {
		ans += "/" + stack[i]
	}
	if ans == "" {
		return "/"
	}
	return ans
}
