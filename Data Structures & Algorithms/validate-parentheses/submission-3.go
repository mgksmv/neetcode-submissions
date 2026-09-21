func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

    stack := []string{}

	for _, ch := range s {
		chStr := string(ch)

		if chStr == "(" || chStr == "{" || chStr == "[" {
			stack = append(stack, string(chStr))
		} else {
			if len(stack) == 0 {
				return false
			}

			topCh := stack[len(stack)-1]
			if (chStr == ")" && topCh != "(") ||
				(chStr == "}" && topCh != "{") ||
				(chStr == "]" && topCh != "[") {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
