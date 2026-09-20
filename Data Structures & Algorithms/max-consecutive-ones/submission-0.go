func findMaxConsecutiveOnes(nums []int) int {
    left := 0
    right := 0

    maxLen := 0

    for right < len(nums) {
        if nums[right] == 0 {
            left = right + 1
        }

        if nums[right] == 1 {
            maxLen = max(maxLen, (right - left) + 1)
        }

        right++
    }

    return maxLen
}
