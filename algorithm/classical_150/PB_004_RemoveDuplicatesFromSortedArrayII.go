package classical_150

// removeDuplicatesII 删除有序数组中的重复项Ⅱ
func removeDuplicatesII(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	cur, size := 0, 1
	for i := 1; i < len(nums); {
		if nums[i] == nums[cur] {
			if size == 2 {
				for i < len(nums) && nums[i] == nums[cur] {
					i++
				}
				if i == len(nums) {
					break
				}
				size = 1
			} else {
				size++
			}
		} else {
			size = 1
		}
		cur++
		nums[cur] = nums[i]
		i++
	}
	return cur + 1
}
