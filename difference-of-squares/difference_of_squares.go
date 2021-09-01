// Package diffsquares contains functions relating to the computation of the
// difference of squares.
package diffsquares

// SquareOfSum returns the square of the sum of consecutive integers up to
// and including max.
func SquareOfSum(max int) int {
	sum := max * (max + 1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of consecutive integers up to
// and including max.
func SumOfSquares(max int) int {
	return max * (max + 1) * (2*max + 1) / 6
}

// Difference returns the difference of the square of sums and the sum of squares,
// up to max each.
func Difference(max int) int {
	return SquareOfSum(max) - SumOfSquares(max)
}
