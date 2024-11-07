package sprint
func FactorialRecursive(n int) int {
	if n < 0 {
		return 0 // Return 0 for non-positive values
	}
	if n > 20 {
		return 0 // Return 0 for potential overflows
	}
	if n == 0 {
		return 1 // Base case: factorial of 0 is 1
	}
	return n * FactorialRecursive(n-1) // Recursive call to calculate factorial
}