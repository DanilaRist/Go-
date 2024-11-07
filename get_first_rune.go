package sprint

func GetFirstRune(s string) rune {
    // Convert the string to a rune slice and return the first element
    return []rune(s)[0]
}