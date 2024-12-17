package classical_150

import "github.com/showyquasar88/proj-combine/algorithm/utils"

// trap 接雨水
func trap(height []int) int {
	stack, idx, res := make([]int, len(height)), -1, 0
	for i := 0; i < len(height); i++ {
		for idx != -1 && height[i] >= height[stack[idx]] {
			cur := stack[idx]
			idx--
			if idx == -1 {
				break
			}
			left := stack[idx]
			curWidth := i - left - 1
			curHeight := utils.Min(height[i], height[left]) - height[cur]
			res += curWidth * curHeight
		}
		idx++
		stack[idx] = i
	}
	return res
}
