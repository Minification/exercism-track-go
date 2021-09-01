// Package raindrops contains the conversion function for raindrop sounds.
package raindrops

import "strconv"

type tuple struct {
	raindrop string
	factor   int
}

var raindropTuples = []tuple{
	tuple{
		raindrop: "Pling", factor: 3,
	}, tuple{
		raindrop: "Plang", factor: 5,
	}, tuple{
		raindrop: "Plong", factor: 7,
	},
}

// Convert returns the raindrop sounds for the specified number.
func Convert(number int) (output string) {

	for _, t := range raindropTuples {
		if number%t.factor == 0 {
			output += t.raindrop
		}
	}

	if len(output) == 0 {
		output = strconv.Itoa(number)
	}
	return
}
