package sprint
func GetLastRune(s string) rune {
    // Convert the string to a rune slice
    runes := []rune(s)
    // Return the last element of the rune slice
    return runes[len(runes)-1]
}