func validPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	str := []rune(s)
	l := 0
	r := len(s) - 1

	for l < r {
		if str[l] != str[r] {
			return is(str, l + 1, r) || is(str, l, r-1)
		}
		l++
		r--
	}
	return true
}

func is(s []rune, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}