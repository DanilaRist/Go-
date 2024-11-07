package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	reversedAlphabet := make([]byte, 0, 26)

	for i := len(alphabet) - 1; i >= 0; i -= step {
		reversedAlphabet = append(reversedAlphabet, alphabet[i])
	}

	return string(reversedAlphabet)
}