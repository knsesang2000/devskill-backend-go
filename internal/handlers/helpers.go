package handlers

// contains checks if needle is in haystack (naive search for MVP)
func contains(haystack, needle string) bool {
	return len(needle) > 0 && len(haystack) >= len(needle) && (func() bool { return indexOf(haystack, needle) >= 0 })()
}

// indexOf returns the first index of sub in s, or -1 if not found.
func indexOf(s, sub string) int {
	n := len(sub)
	if n == 0 {
		return 0
	}
	for i := 0; i+n <= len(s); i++ {
		if s[i:i+n] == sub {
			return i
		}
	}
	return -1
}
