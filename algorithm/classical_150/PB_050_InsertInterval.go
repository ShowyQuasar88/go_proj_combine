package classical_150

import "slices"

// insertInterval 插入区间
func insertInterval(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	begin, end, ans := intervals[0][0], intervals[0][1], make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if end < intervals[i][0] {
			ans = append(ans, []int{begin, end})
			begin, end = intervals[i][0], intervals[i][1]
		} else if end < intervals[i][1] {
			end = intervals[i][1]
		}
	}
	ans = append(ans, []int{begin, end})
	return ans
}
