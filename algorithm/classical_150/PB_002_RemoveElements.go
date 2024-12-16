package classical_150

// removeElement 移除元素
func removeElement(nums []int, val int) int {
	last := len(nums)
	for i := 0; i < last; {
		if nums[i] == val {
			last--
			nums[i] = nums[last]
		} else {
			i++
		}
	}
	return last
}
