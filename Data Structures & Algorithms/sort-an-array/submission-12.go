func sortArray(nums []int) []int {
	if len(nums) > 0 {
		quickSort(nums, 0, len(nums) - 1)
	}
	return nums
}

func quickSort(nums []int, low, hight int) {
	if low < hight {
		pivot := partion(nums, low, hight)

		quickSort(nums, low, pivot - 1)
		quickSort(nums, pivot + 1, hight)
	}
}

func partion(nums []int, low, hight int) int{
	pivot := nums[hight]
	
	i := low - 1
	for j := low; j < hight; j++ {
		if nums[j] <= pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	} 

	nums[i + 1], nums[hight] = nums[hight], nums[i + 1]
	return i + 1
}
