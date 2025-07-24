package intersectionArray

func IntersectionArray(num1 []int, num2 []int) []int {
	set := make(map[int]struct{})
	for _, v := range num1 {
		set[v] = struct{}{}
	}

	resultSet := make(map[int]struct{})
	for _, v := range num2 {
		if _, exist := set[v]; exist {
			resultSet[v] = struct{}{}
		}
	}

	result := make([]int, 0, len(resultSet))
	for v := range resultSet {
		result = append(result, v)
	}

	return result
}
