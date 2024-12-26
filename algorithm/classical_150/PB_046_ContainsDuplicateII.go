package classical_150

// containsNearbyDuplicate 存在重复元素 II
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int, len(nums))
	for idx, v := range nums {
		if mpv, ok := mp[v]; ok && idx-mpv <= k {
			return true
		}
		mp[v] = idx
	}
	return false
}
