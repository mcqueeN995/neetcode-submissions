import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := make([][]int, 0, 0)
	s := 0

	for i := 0; i < len(nums); i++{
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l, r := i + 1, len(nums)-1
		for l < r {
			s = nums[i] + nums[r] + nums[l]
			if s < 0 {
				l++
			} else if s > 0 {
				r--
			}  else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				r--
				l++
				for l < r && nums[l] == nums[l - 1] {
					l++
				}

				for l < r && nums[r] == nums[r + 1]{
					r--
				}
			}

			
		}
	}	
	return ans
}
