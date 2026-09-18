func twoSum(nums []int, target int) []int {
    numIndexes := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if index, exists := numIndexes[diff]; exists {
			return []int{index, i}
		}
		numIndexes[nums[i]] = i 
	}

	return []int{}
}
