package averageOfEvenNumbers

func AverageOfEvenNumbers(nums []int) float64 {
	sum := 0.0
	countOfEvenNumbers := 0

	for _, num := range nums {
		if num%2 == 0 {
			sum += float64(num)
			countOfEvenNumbers++
		}
	}
	if countOfEvenNumbers == 0 {
		return 0.0
	}
	return sum / float64(countOfEvenNumbers)
}
