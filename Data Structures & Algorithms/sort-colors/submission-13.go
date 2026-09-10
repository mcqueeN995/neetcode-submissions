func sortColors(nums []int) {
    w1, w2 := 0, len(nums) - 1
	r := 0

	for r <= w2{
		if nums[r] == 0 { 
			nums[r], nums[w1] = nums[w1], nums[r]
			w1++
		} else if nums[r] == 2 { 
			nums[r], nums[w2] = nums[w2], nums[r]
			w2--
			continue
		}
		r++
	}
}


