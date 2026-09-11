func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	ans := make([]int, 0, 0)
	m := make(map[int]int)

	for _, elm := range nums{
		m[elm]++
	}

	for key, val := range m {
		if val > len(nums) / 3{
			ans = append(ans, key)
		}
	}

	return ans
}
