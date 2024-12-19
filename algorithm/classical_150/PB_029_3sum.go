package classical_150

import "sort"

// threeSum 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := len(nums) - 1; i > 1; i-- {
		for i > 1 && i != len(nums)-1 && nums[i] == nums[i+1] {
			i--
		}
		tmp := twoSum2(nums, 0, i-1, -nums[i])
		for _, item := range tmp {
			item = append(item, nums[i])
			ans = append(ans, item)
		}
	}
	return ans
}

func twoSum2(nums []int, begin, end, target int) [][]int {
	res, left, right := make([][]int, 0), begin, end
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			res = append(res, []int{nums[left], nums[right]})
			left++
			right--
		}
		for right > 0 && right < end && nums[right] == nums[right+1] {
			right--
		}
		for left < len(nums) && left > begin && nums[left] == nums[left-1] {
			left++
		}
	}
	return res
}
