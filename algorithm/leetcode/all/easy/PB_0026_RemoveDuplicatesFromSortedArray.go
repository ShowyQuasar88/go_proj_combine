package easy

// removeDuplicates 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	curIdx := 0
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] != nums[curIdx] {
			curIdx++
			nums[curIdx] = nums[idx]
		}
	}
	return curIdx + 1
}
