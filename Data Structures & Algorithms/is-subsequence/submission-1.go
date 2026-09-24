func isSubsequence(s string, t string) bool {
    s_p := 0
	t_p := 0

	for s_p < len(s) && t_p < len(t) {
		if s[s_p] == t[t_p] {
			s_p++
		}
		t_p++
	}

	return s_p == len(s)
}
