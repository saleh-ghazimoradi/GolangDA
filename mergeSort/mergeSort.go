package mergeSort

// MergeSort sorts the input slice using the merge sort algorithm
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Divide the array into two halves
	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	// Recursively sort both halves
	left = MergeSort(left)
	right = MergeSort(right)

	// Merge the sorted halves
	return merge(left, right)
}

// merge combines two sorted slices into a single sorted slice
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	leftIdx, rightIdx := 0, 0

	// Compare elements from both slices and merge in sorted order
	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] <= right[rightIdx] {
			result = append(result, left[leftIdx])
			leftIdx++
		} else {
			result = append(result, right[rightIdx])
			rightIdx++
		}
	}

	// Append remaining elements from left slice, if any
	result = append(result, left[leftIdx:]...)

	// Append remaining elements from right slice, if any
	result = append(result, right[rightIdx:]...)

	return result
}
