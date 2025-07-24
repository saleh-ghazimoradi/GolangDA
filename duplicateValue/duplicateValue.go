package duplicateValue

func DuplicateValue(letters []string) string {
	seen := make(map[string]struct{})
	for _, v := range letters {
		if _, exists := seen[v]; exists {
			return v
		}
		seen[v] = struct{}{}
	}
	return ""
}
