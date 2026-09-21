func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix) - 1

	for left <= right {
		mid := left + (right - left) / 2
		row := matrix[mid]

		if target < row[0] {
			right = mid - 1
		} else if target > row[len(row) - 1] {
			left = mid + 1
		} else {
			return searchInsideRow(row, target)
		}
	}

	return false
}

func searchInsideRow(row []int, target int) bool {
	left := 0
	right := len(row) - 1

	for left <= right {
		mid := left + (right - left) / 2
		guess := row[mid]

		if guess == target {
			return true
		}

		if guess < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
