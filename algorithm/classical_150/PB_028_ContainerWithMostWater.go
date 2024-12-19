package classical_150

// maxArea 盛水最多的容器
func maxArea(height []int) int {
	left, right, ans := 0, len(height)-1, 0
	for left < right {
		cur := (right - left)
		if height[left] > height[right] {
			cur *= height[right]
			right--
		} else {
			cur *= height[left]
			left++
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
