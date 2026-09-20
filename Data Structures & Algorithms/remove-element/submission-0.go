func removeElement(nums []int, val int) int {
    left := 0
	right := len(nums)

	for left < right {
		if nums[left] == val {
			right--
			nums[left] = nums[right]
		} else {
			left++
		}
	}

	return right
}
