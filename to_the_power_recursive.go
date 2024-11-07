package sprint
func ToThePowerRecursive(n int, power int) int {
	if power < 0 {
		return 0 // Return 0 for negative powers
	}
	if power == 0 {
		return 1 // Base case: n^0 is 1
	}
	return n * ToThePowerRecursive(n, power-1) // Recursive call to calculate power
}
