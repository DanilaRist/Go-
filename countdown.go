package sprint

func Countdown(n int) string {
	if n < 0 || n > 9 {
		return "" // Return an empty string for invalid inputs
	}

	var countdown string

	for i := n; i > 0; i -= 2 {
		countdown += string('0' + i) + ", "
	}

	return countdown + "0!"
}