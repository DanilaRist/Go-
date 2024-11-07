package sprint
func StrLength(s string) []int {
    // Initialize counters for runes and bytes
    runeCount := 0
    byteLength := 0

    // Iterate over the string
    for _, char := range s {
        // Increment the rune count
        runeCount++
        // Increment the byte length by the size of the current rune
        byteLength += runeByteLength(char)
    }

    // Return the results as a slice of integers
    return []int{runeCount, byteLength}
}

// Function to calculate byte length of a rune
func runeByteLength(r rune) int {
    switch {
    case r < 128:
        return 1
    case r < 2048:
        return 2
    case r < 65536:
        return 3
    default:
        return 4
    }
}
