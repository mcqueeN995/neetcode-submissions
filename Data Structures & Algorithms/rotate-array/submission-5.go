func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return 
	}

	l, r := 0, len(nums) - 1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	l, r = 0, k - 1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	l, r = k, len(nums)-1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
