package isPalindrome

func IsPalindrome(s string) bool {
	leftIndex, rightIndex := 0, len(s)-1

	for leftIndex < len(s)/2 {
		if s[leftIndex] != s[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}
	return true
}
