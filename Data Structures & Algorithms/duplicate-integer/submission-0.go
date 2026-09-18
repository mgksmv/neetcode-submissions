func hasDuplicate(nums []int) bool {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if i != j && nums[i] == nums[j] {
                return true
            }
        }
    }

    return false
}
