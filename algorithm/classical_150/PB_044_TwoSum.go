package classical_150

// twoSum 两数之和
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for idx, v := range nums {
		_, ok := mp[target-v]
		if ok {
			return []int{mp[target-v], idx}
		}
		mp[v] = idx
	}
	return nil
}
