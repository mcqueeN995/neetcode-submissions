func checkInclusion(s1 string, s2 string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s1); i++{
		m[s1[i]]++
	}

	left := 0
	for right := 0; right < len(s2); right++ {
		m[s2[right]]--

		if right - left + 1 > len(s1){
			m[s2[left]]++
			left++
		}

		if allZero(m) {
			return true
		}
	}

	return false
}


func allZero(m map[byte]int) bool {
	for _, val := range m{
		if val != 0 {
			return false
		}
	}
	return true
}