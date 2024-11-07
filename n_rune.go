package sprint

func NRune(s string, i int) rune {
    // Convert the string to a rune slice
    runes := []rune(s)
    // Return the rune at the specified index
    return runes[i]
}
