package greatesProductOfAnyThreeNumbers

import "sort"

func MaxProductOfThree(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1] * nums[n-2] * nums[n-3]
}
