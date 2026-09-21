func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
