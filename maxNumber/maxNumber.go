package maxNumber

import "sort"

func QuadraticMaxNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxNum := nums[0]
	for i := 0; i < len(nums); i++ {
		isMax := true
		for j := 0; j < len(nums); j++ {
			if nums[j] > nums[i] {
				isMax = false
				break
			}
		}
		if isMax {
			maxNum = nums[i]
		}
	}
	return maxNum
}

func MaxNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func LinearMaxNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}
