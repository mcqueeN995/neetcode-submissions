func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	for _, char := range t {
		m[char]--
	}

	for _, freq := range m {
		if freq != 0 {
			return false
		}
	}
	return true
}
