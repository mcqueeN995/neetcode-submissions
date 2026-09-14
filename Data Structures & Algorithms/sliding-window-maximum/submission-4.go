func maxSlidingWindow(nums []int, k int) []int {
    deqeue := make([]int, 0)
	result := make([]int, 0)

	for i := 0; i < len(nums); i++{
		if len(deqeue) > 0 && deqeue[0] <= i - k {
			deqeue = deqeue[1:]
		}

		for len(deqeue) > 0 && nums[deqeue[len(deqeue) - 1]] <= nums[i] {
			 deqeue = deqeue[:len(deqeue)-1]
		}

		deqeue = append(deqeue, i)
		if i >= k - 1{
			result = append(result, nums[deqeue[0]])
		}
	}
	return result
}
