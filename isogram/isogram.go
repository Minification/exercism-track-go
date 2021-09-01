// Package isogram contains functions to check if a word is an isogram.
package isogram

import "unicode"

// IsIsogram checks if the word is an isogram.
func IsIsogram(word string) bool {
	var seen [26]bool
	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}
		// Map runes to alphabet indices
		var i = unicode.ToLower(char) - 'a'
		if seen[i] {
			return false
		}
		seen[i] = true
	}
	return true
}
