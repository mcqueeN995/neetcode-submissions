func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	s, l, r := 0, 0 ,len(numbers) - 1

	for l < r {
		s = numbers[l] + numbers[r]

		if s < target {
			l++
		} else if s > target {
			r--
		} else {
			return []int{l+1, r+1}
		}
	}
	return []int{0, 0}
}
