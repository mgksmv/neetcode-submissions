func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sFreqs := make(map[byte]int)
	tFreqs := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sFreqs[s[i]]++
		tFreqs[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if sFreqs[s[i]] != tFreqs[s[i]] {
			return false
		}
	}

	return true
}
