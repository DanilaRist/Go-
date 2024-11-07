package sprint 
func FactorialIterative(n int) int {
	// Handle non-positive values
	if n < 0 {
		return 0
	}

	// Handle overflow for large factorials
	if n > 20 {
		// Overflow warning (without fmt)
	}

	result := 1
	for i := 1; i <= n; i++ {
		// Check for overflow
		if result > (1<<31-1)/i {
			// Overflow detected (without fmt)
			return 0
		}
		result *= i
	}

	return result
}