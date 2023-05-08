package main

// Complete the solution so that it splits the string into pairs of two characters.
// If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
func Solution(str string) []string {
	runes := []rune(str)
	n := len(runes)
	var pairs []string
	for i := 0; i < n; i += 2 {
		if i+1 < n {
			pairs = append(pairs, string(runes[i])+string(runes[i+1]))
		} else {
			pairs = append(pairs, string(runes[i])+"_")
		}
	}
	return pairs
}
