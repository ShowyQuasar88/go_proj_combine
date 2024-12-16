package classical_150

// canJump 跳跃游戏
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxRight := nums[0]
	for i := 1; i <= maxRight; i++ {
		if i+nums[i] > maxRight {
			maxRight = i + nums[i]
		}
		if maxRight >= len(nums)-1 {
			return true
		}
	}
	return false
}
