package diffsquares

// SquareOfSum sums all the natural numbers up to `n`, inclusive,
// and returns the square of that sum.
func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares returns the sum of the squares of all the
// natural numbers up to `n`, inclusive.
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between the square of sums up to `n`
// and the sum of squares up to `n`
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
