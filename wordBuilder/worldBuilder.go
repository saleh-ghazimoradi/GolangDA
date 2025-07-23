package wordBuilder

func WordBuilder(arr []string) []string {
	collection := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				collection = append(collection, arr[i]+arr[j])
			}
		}
	}
	return collection
}
