package sprint
func ConvertAnyToDec(s string, base string) int {
    // Validate the base
    if len(base) < 2 {
        return 0
    }
    for _, ch := range base {
        if ch == '+' || ch == '-' {
            return 0
        }
    }

    // Create a map to store the index of each character in the base string
    indexMap := make(map[rune]int)
    for i, ch := range base {
        indexMap[ch] = i
    }

    // Initialize the result to 0
    result := 0

    // Iterate over each digit in the numeric string and calculate its value in decimal
    for _, digit := range s {
        index, ok := indexMap[digit]
        if !ok {
            return 0 // The digit is not in the base, return 0
        }
        result = result*len(base) + index
    }

    return result
}
