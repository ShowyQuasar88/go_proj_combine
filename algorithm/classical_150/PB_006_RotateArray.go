package classical_150

// rotate 轮转数组
func rotate(nums []int, k int) {
	length := len(nums)
	if k == 0 || (k >= length && length%k == 0) {
		return
	}
	k = k % length
	res := make([]bool, len(nums))
	for i := 0; i < length; i++ {
		if res[i] {
			continue
		}
		prev, tmp := nums[i], 0
		for j := (i + k) % length; !res[j]; j = (j + k) % length {
			tmp = nums[j]
			nums[j] = prev
			prev = tmp
			res[j] = true
		}
	}
}
