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

	if len(sFreqs) != len(tFreqs) {
		return false
	}

	for l, v := range sFreqs {
		if tFreqs[l] != v {
			return false
		}
	}

	return true
}
