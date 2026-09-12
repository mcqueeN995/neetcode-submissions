func characterReplacement(s string, k int) int {
	l := 0
	m := make(map[byte]int)
	ans := 0
	maxCount := 0

	for r := 0; r < len(s); r++{
		charR := s[r]
		charL := s[l]
		m[charR]++

		if maxCount < m[charR] {
			maxCount = m[charR]
		}

		size := r - l + 1
		for size - maxCount > k {
			m[charL]--
			l++
			size = r - l + 1
		}

		if size > ans {
			ans = size
		}

	}
	return ans
	
}





