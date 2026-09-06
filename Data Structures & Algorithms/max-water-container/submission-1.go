func maxArea(heights []int) int {
	if len(heights) <= 1 {
		return 0
	}

	l, r := 0, len(heights) - 1
	ans := 0

	for l < r {
		c := (r - l) * min(heights[l], heights[r])  

		if c > ans {
			ans = c
		}

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}
