package sprint

func containsChar(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

func containsDuplicateChars(s string) bool {
	seen := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			return true
		}
		seen[s[i]] = true
	}
	return false
}

func NbrBase(n int, base string) string {
	// Validate the base
	if len(base) < 2 || containsDuplicateChars(base) || containsChar(base, '+') || containsChar(base, '-') {
		return "NV"
	}

	// Handle the sign separately
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	// Convert n to the specified base
	baseLength := len(base)
	result := ""
	for n > 0 {
		remainder := n % baseLength
		result = string(base[remainder]) + result
		n /= baseLength
	}

	// If result is empty, it means n was 0
	if result == "" {
		result = string(base[0])
	}

	return sign + result
}
