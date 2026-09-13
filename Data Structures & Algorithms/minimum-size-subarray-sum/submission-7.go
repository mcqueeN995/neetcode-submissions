
func minSubArrayLen(target int, nums []int) int {
	sum, left, c, ans := 0, 0, 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		c++

		for sum >= target {
			if c < ans || ans == 0{
				ans = c
			}
			c--
			sum-=nums[left]
			left++
		}
	}
	
	return ans
}



