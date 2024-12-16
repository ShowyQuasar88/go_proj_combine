package classical_150

// jump 跳跃游戏II
func jump(nums []int) int {
	if len(nums) == 1 || (len(nums) > 1 && nums[0] == 0) {
		return 0
	}
	next, maxRight, cnt, length := 0, 0, 0, len(nums)
	for i := 0; i <= maxRight; i++ {
		if i+nums[i] > next {
			next = i + nums[i]
		}
		if i == maxRight && maxRight < length-1 {
			cnt++
			maxRight = next
			if maxRight >= length-1 {
				return cnt
			}
		}
	}
	return cnt
}
