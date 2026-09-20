func replaceElements(arr []int) []int {
	greatest := -1

    for i := len(arr) - 1; i >= 0; i-- {
        current := arr[i]
        arr[i] = greatest

        if current > greatest {
            greatest = current
        }
    }

    return arr
}
