import ma "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	
	m := make(map[rune]int)
	for _, char := range t {
		m[rune(char)]++
	}
	
	have, need := 0, len(m)
	resLen := ma.MaxInt32
	l := 0
	window := make(map[rune]int)
	res := []int{-1, -1}

	for r := 0; r < len(s); r++{
		char := rune(s[r])
		window[char]++

		if m[char] > 0 && m[char] == window[char] {
			have++
		}

		for have == need {
			if r - l + 1 < resLen{
			res = []int{l, r}
			resLen = r - l + 1
			}

			window[rune(s[l])]--
			if m[rune(s[l])] > 0 && window[rune(s[l])] < m[rune(s[l])] {
				have--
			}
			l++
		}
	}

	if res[0] == -1 {
		return ""
	}
	return s[res[0]:res[1]+1]

}
