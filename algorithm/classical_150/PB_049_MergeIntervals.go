package classical_150

import (
	"slices"
)

// mergeIntervals 合并区间
func mergeIntervals(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	begin, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ans = append(ans, []int{begin, end})
			begin, end = intervals[i][0], intervals[i][1]
		} else if intervals[i][1] > end {
			end = intervals[i][1]
		}
	}
	ans = append(ans, []int{begin, end})
	return ans
}
