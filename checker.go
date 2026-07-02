package main

func findClosest(word string, dictionary []string) string {
	for _, dictWord := range dictionary {
		if isSimilar(word, dictionary) {
			return dictWord
		}
	}
	return ""
}
