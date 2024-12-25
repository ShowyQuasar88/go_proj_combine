package classical_150

// twoSum 两数之和 II - 输入有序数组
func twoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{0, 0}
}
