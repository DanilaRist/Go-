package sprint
func NthFibonacci(index int) int {
	if index < 0 {
		return -1 // Return -1 for negative indices
	}
	if index == 0 {
		return 0 // Base case: Fibonacci(0) is 0
	}
	if index == 1 {
		return 1 // Base case: Fibonacci(1) is 1
	}
	return NthFibonacci(index-1) + NthFibonacci(index-2) // Recursive call to calculate Fibonacci sequence
}
