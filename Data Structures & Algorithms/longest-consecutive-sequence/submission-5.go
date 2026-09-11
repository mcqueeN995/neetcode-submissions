func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool)
	ans := 1

	for _, elm := range nums {
		m[elm] = true
	}

	for _, elm := range nums {
		if m[elm - 1] == false {
			key :=  elm
			c := 1
			for m[key + 1] == true {
				c++
				key++
			}
			if c > ans {
				ans = c
			}
		}
		
	}
	return ans
}	

