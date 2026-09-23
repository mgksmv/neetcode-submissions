func removeDuplicates(nums []int) int {
	left := 0
	right := 0

	for right < len(nums) {
		nums[left] = nums[right]

		for right < len(nums) && nums[left] == nums[right] {
			right++
		}

		left++
	}

	return left
}
