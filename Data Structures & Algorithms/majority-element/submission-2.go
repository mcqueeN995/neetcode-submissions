func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, elm := range nums{
		m[elm]++
	}

	for key, val := range m {
		if val > len(nums)/2 {
			return key
		}
	}
	return 0
}
