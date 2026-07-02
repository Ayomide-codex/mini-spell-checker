package main

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := []string{"hello", "world", "golang", "project"}
	input := "helo word"

	words := strings.Fields(input)
	fmt.Println("Did you mean:")

	for _, word := range words {
		suggestion := findClosest(word, dictionary)
		if suggestion != "" {
			fmt.Println(suggestion)
		}
	}
}
