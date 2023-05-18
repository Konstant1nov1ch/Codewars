package main

import (
	"strings"
	"unicode"
)

func toWeirdCase(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		words[i] = applyWeirdCase(word)
	}
	return strings.Join(words, " ")
}

func applyWeirdCase(word string) string {
	result := []rune(word)
	for i := range result {
		if i%2 == 0 {
			result[i] = unicode.ToUpper(result[i])
		} else {
			result[i] = unicode.ToLower(result[i])
		}
	}
	return string(result)
}
