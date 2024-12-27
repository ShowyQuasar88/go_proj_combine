package classical_150

import "strconv"

// evalRPN 逆波兰表达式求值
func evalRPN(tokens []string) int {
	nums := make([]int, 0, len(tokens)/2)
	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err != nil {
			left, right := nums[len(nums)-2], nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			switch token {
			case "+":
				nums[len(nums)-1] = left + right
			case "-":
				nums[len(nums)-1] = left - right
			case "*":
				nums[len(nums)-1] = left * right
			case "/":
				nums[len(nums)-1] = left / right
			}
		} else {
			nums = append(nums, v)
		}
	}
	return nums[0]
}
