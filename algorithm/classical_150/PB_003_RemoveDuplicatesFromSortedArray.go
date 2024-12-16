package classical_150

// removeDuplicates 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	cur := 0
	for i := 1; i < len(nums); {
		for i < len(nums) && nums[i] == nums[cur] {
			i++
		}
		if i == len(nums) {
			break
		}
		cur++
		nums[cur] = nums[i]
		i++
	}
	return cur + 1
}
