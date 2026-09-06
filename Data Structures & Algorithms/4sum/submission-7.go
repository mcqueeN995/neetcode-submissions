import "slices"

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	ans := make([][]int, 0, 0)
	s := 0

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i - 1]{
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if  j > i + 1 && nums[j] == nums[j - 1]{
				continue
			}

			l, r := j + 1, len(nums) - 1

			for l < r {
				s = nums[i] + nums[j] + nums[l] + nums[r]

				if s < target {
					l++
				} else if s > target {
					r--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--

					for l < r && nums[l] == nums[l - 1] {
						l++
					}

					for l < r && nums[r] == nums[r + 1] {
						r--
					}
				}
			}
		}
	}
	return ans
}
