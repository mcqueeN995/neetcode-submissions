func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for idx, elm := range nums{
		if i, found := m[target - elm]; found{
			return []int{i, idx}
		}
		m[elm] = idx
	}

	return []int{0, 0}
}
