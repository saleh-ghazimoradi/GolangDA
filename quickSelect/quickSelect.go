package quickSelect

func QuickSelect(arr []int, k int) int {
	result := make([]int, len(arr))
	copy(result, arr)

	return quickSelect(result, 0, len(result)-1, k-1)
}

func quickSelect(arr []int, low, high, k int) int {
	if low == high {
		return arr[low]
	}

	pi := partition(arr, low, high)

	if pi == k {
		return arr[pi]
	}

	if k < pi {
		return quickSelect(arr, low, pi-1, k)
	}

	return quickSelect(arr, pi+1, high, k)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
