// Package hamming contains a function to calculate the Hamming distance between two strings.
// Adapted from bobahop's solution:
// https://exercism.io/tracks/go/exercises/hamming/solutions/2f12fec07a3d46309a2c2fbdf42543f4
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates the Hamming distance between two strings of equal length.
func Distance(a, b string) (distance int, err error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("strands must be of equal length")
	}
	aRunes := []rune(a)
	bRunes := []rune(b)
	for i, r := range aRunes {
		if r != bRunes[i] {
			distance++
		}
		i++
	}
	return distance, nil
}
