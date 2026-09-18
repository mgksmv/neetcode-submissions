func hasDuplicate(nums []int) bool {
    visitedNums := make(map[int]bool)

    for _, num := range nums {
        if _, exists := visitedNums[num]; exists {
            return true
        }

        visitedNums[num] = true
    }

    return false
}
