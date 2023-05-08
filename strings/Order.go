package main

//Your task is to sort a given string. Each word in the string will contain a single number. This number is the position the word should have in the result.
//
//Note: Numbers can be from 1 to 9. So 1 will be the first word (not 0).
//
//If the input string is empty, return an empty string. The words in the input String will only contain valid consecutive numbers.

import (
	"strconv"
	"strings"
)

func Order(sentence string) string {
	words := strings.Split(sentence, " ")
	sortedWords := make([]string, len(words))
	for _, word := range words {
		for _, char := range word {
			if n, err := strconv.Atoi(string(char)); err == nil {
				sortedWords[n-1] = word
				break
			}
		}
	}
	return strings.Join(sortedWords, " ")
}
