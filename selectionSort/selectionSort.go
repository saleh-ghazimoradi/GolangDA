package selectionSort

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] < arr[j] {
				minIndex = j
			}
			arr[j], arr[minIndex] = arr[minIndex], arr[j]
		}
	}
	return arr
}
