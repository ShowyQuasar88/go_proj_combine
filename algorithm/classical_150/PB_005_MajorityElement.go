package classical_150

// majorityElement 多数元素
func majorityElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	cur, hp := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == cur {
			hp++
		} else if hp == 0 {
			cur = nums[i]
			hp = 1
		} else {
			hp--
		}
	}
	return cur
}
