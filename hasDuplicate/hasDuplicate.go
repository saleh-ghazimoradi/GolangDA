package hasDuplicate

// HasDuplicate is less efficient
func HasDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// EfficientHasDuplicate is more efficient
func EfficientHasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
