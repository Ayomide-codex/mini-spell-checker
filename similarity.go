package main

func isSimilar(a, b string) bool {
	// if lengths differ by more than 2 they can't be similar
	diff := len(a) - len(b)
	if diff < 0 {
		diff = -diff
	}
	if diff > 2 {
		return false
	}

	// count how many characters match
	matches := 0
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			matches++
		}
	}

	// similar if more than half the characters match
	return matches >= len(b)-2
}
