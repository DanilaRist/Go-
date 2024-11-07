package sprint
func StrReverse(s string) string {
    // Convert the string to a rune slice
    runes := []rune(s)
    
    // Initialize indices for reversing
    start := 0
    end := len(runes) - 1
    
    // Swap characters from start to end
    for start < end {
        runes[start], runes[end] = runes[end], runes[start]
        start++
        end--
    }
    
    // Convert the reversed rune slice back to a string
    return string(runes)
}
