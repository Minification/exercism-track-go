// Package luhn contains functions to check if a number is valid
// as per the Luhn formula.
package luhn

var lookup = [2][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{0, 2, 4, 6, 8, 1, 3, 5, 7, 9},
}

// Valid checks whether the specified string represents a valid number according
// to the Luhn formula.
func Valid(number string) bool {
	var i, sum, digitCount, double int
	for i = len(number) - 1; i >= 0; i-- {
		digit := number[i]
		switch {
		case '0' <= digit && digit <= '9':
			sum += lookup[double][int(digit-'0')]
			double ^= 1
			digitCount++
		default:
			if digit == ' ' {
				continue
			}
			return false
		}
	}
	return sum%10 == 0 && digitCount > 1
}
