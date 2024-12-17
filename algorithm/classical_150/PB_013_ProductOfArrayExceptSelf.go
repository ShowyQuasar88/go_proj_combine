package classical_150

// productExceptSelf 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	left, right, length := 1, 1, len(nums)
	for i := 1; i < length; i++ {
		left *= nums[i-1]
		res[i] = left
	}
	for i := length - 2; i >= 0; i-- {
		right *= nums[i+1]
		res[i] *= right
	}
	res[0] = right
	return res
}
