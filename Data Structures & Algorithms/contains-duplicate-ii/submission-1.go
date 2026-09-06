func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]bool)

	for r := 0; r < len(nums); r++{
		if (window[nums[r]]) {
			return true
		}

		window[nums[r]] = true

		if r >= k {
			delete(window, nums[r - k])
		}
	}
	return false
}
