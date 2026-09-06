func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	r1 := []rune(s)
	r2 := []rune(t)
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})
	return string(r1) == string(r2)
}
