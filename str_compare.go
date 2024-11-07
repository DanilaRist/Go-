package sprint
func StrCompare(a, b string) int {
    minLength := len(a)
    if len(b) < minLength {
        minLength = len(b)
    }

    // Iterate over each character in the strings
    for i := 0; i < minLength; i++ {
        if a[i] < b[i] {
            return -1
        } else if a[i] > b[i] {
            return 1
        }
    }

    // If all characters are equal up to the length of the shorter string, compare string lengths
    if len(a) < len(b) {
        return -1
    } else if len(a) > len(b) {
        return 1
    }

    // Strings are equal
    return 0
}
