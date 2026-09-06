func mergeAlternately(word1 string, word2 string) string {
	if len(word1) == 0{
		return word2
	}

	if len(word2) == 0{
		return word1
	}

	p1, p2 := 0, 0
	s := ""

	for p1 < len(word1) && p2 < len(word2) {
		s += string(word1[p1]) + string(word2[p2])
		p1++
		p2++
	}

	//aba
	//zxcc
	//azbxac

	// len(aba) = 3 p1 = 3
	// len(zxcc) = 4 p2 = 4


	
	return s + word1[p1:] + word2[p2:]
}
