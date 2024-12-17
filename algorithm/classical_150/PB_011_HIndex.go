package classical_150

import "sort"

// hIndex H指数
func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
		} else {
			break
		}
	}
	return h
}
