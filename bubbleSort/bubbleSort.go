package bubbleSort

func BubbleSort(arr []int) []int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return sorted
}
