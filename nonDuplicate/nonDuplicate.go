package nonDuplicate

func NonDuplicate(s string) rune {
	freq := make(map[rune]int)

	for _, char := range s {
		freq[char]++
	}

	for _, char := range s {
		if freq[char] == 1 {
			return char
		}
	}

	return 0
}
