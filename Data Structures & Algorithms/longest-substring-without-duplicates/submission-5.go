func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	l := 0
	res := 0
	rs := []rune(s)
	window := make(map[rune]int)

	for r := 0; r < len(rs); r++ {
		char := rs[r]

		if idx, exist := window[char]; exist && idx >= l {
			l = idx + 1
		}

		window[char] = r

		if res < r - l + 1 {
			res = r - l + 1
		}
	}
	return res
}
