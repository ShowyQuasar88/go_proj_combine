package classical_150

import (
	"fmt"
	"strconv"
)

// summaryRanges 汇总区间
func summaryRanges(nums []int) []string {
	length, ans := len(nums), make([]string, 0)
	if length == 0 {
		return ans
	}
	if length == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	begin, end := nums[0], nums[0]
	for i := 1; i < length; i++ {
		if nums[i-1]+1 < nums[i] {
			if begin == end {
				ans = append(ans, strconv.Itoa(begin))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", begin, end))
			}
			begin, end = nums[i], nums[i]
		} else {
			end = nums[i]
		}
	}
	if begin == end {
		ans = append(ans, strconv.Itoa(begin))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", begin, end))
	}
	return ans
}
