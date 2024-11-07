package sprint

func ReverseAlphabetValue(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		distance := ch - 'a'
		inverse := 'z' - distance

		return inverse
	} else {
		return ch
	}
}