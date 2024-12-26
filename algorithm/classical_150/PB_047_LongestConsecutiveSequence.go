package classical_150

// longestConsecutive 最长连续序列
func longestConsecutive(nums []int) int {
	mp, longest := make(map[int]bool), 0
	for _, v := range nums {
		mp[v] = true
	}
	for num := range mp {
		if !mp[num-1] {
			curLongest, curNum := 1, num+1
			for mp[curNum] {
				curLongest++
				curNum++
			}
			if curLongest > longest {
				longest = curLongest
			}
		}
	}
	return longest
}
