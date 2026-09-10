func sortArray(nums []int) []int {
    if len(nums) > 0 {
        quikSort(nums, 0, len(nums)-1)
    }
    return nums
}


func quikSort(arr []int, low, high int) {
	if low < high {
		pivot := partion(arr, low, high)
		quikSort(arr, pivot+1, high)
		quikSort(arr, low, pivot-1)
	}
}

func partion(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1
	for j := low; j < high; j++{
		if arr[j] <= pivot{
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i + 1], arr[high] = arr[high], arr[i + 1]
	return i + 1
}