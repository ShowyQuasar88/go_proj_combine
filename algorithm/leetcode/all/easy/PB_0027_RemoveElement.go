package easy

// removeElement 移除元素
func removeElement(nums []int, val int) int {
	curIdx := 0
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] != val {
			nums[curIdx] = nums[idx]
			curIdx++
		}
	}
	return curIdx
}
