func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    maxCount := 0

    for _, num := range nums {
        if num == 1 {
            count++
            maxCount = max(maxCount, count)
        } else {
            count = 0
        }
    }

    return maxCount
}
