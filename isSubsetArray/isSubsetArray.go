package isSubsetArray

// IsSubsetArray is less efficient
func IsSubsetArray(slice1 []int, slice2 []int) bool {
	smallerSlice := make([]int, 0)
	largerSlice := make([]int, 0)

	if len(slice1) < len(slice2) {
		smallerSlice = slice1
		largerSlice = slice2
	} else {
		smallerSlice = slice2
		largerSlice = slice1
	}

	for _, val := range smallerSlice {
		foundMatch := false

		for _, largerVal := range largerSlice {
			if val == largerVal {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			return false
		}
	}
	return true
}

// MoreEfficientIsSubsetArray is more efficient using map
func MoreEfficientIsSubsetArray(slice1 []int, slice2 []int) bool {
	largerSlice, smallerSlice := make([]int, 0), make([]int, 0)

	if len(slice1) < len(slice2) {
		smallerSlice = slice1
		largerSlice = slice2
	} else {
		smallerSlice = slice2
		largerSlice = slice1
	}

	set := make(map[int]struct{})
	for _, val := range largerSlice {
		set[val] = struct{}{}
	}
	for _, val := range smallerSlice {
		if _, exists := set[val]; !exists {
			return false
		}
	}
	return true
}
