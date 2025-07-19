package greatesValue

func GreatestValue(values ...int) int {
	greatest := 0
	for _, value := range values {
		if value > greatest {
			greatest = value
		}
	}
	return greatest
}
