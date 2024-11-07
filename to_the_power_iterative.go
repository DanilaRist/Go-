package sprint
func ToThePowerIterative(n int, power int) int {
	if power < 0 {
		return 0 // Return 0 for negative powers
	}

	result := 1
	for i := 0; i < power; i++ {
		result *= n
	}
	return result
}
