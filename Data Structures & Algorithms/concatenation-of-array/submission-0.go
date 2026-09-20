func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2 * n)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		ans[i] = num
		ans[i + n] = num
	}

	return ans
}
