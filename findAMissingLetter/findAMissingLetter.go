package findAMissingLetter

import "strings"

func FindMissingLetter(s string) rune {
	s = strings.ToLower(s)
	seen := [26]bool{}

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			seen[char-'a'] = true
		}
	}

	for i := 0; i < 26; i++ {
		if !seen[i] {
			return rune('a' + i)
		}
	}

	return 0
}
