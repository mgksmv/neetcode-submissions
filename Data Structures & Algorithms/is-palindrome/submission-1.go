func isAlphaNumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for !isAlphaNumeric(s[left]) && left < right {
			left++
		}
		for !isAlphaNumeric(s[right]) && left < right {
			right--
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}
