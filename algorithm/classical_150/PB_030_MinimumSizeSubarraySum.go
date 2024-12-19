package classical_150

// minSubArrayLen 长度最小的子数组
func MinSubArrayLen(target int, nums []int) int {
	left, right, cur, ans := 0, 1, nums[0], 100001
	for left < right && right <= len(nums) {
		if cur < target {
			if right == len(nums) {
				break
			}
			cur += nums[right]
			right++
			continue
		}
		if cur == target && right-left < ans {
			ans = right - left
		}
		cur -= nums[left]
		left++

	}
	if ans == 100001 {
		return 0
	}
	return ans
}
