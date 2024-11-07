package sprint
func ToUpperCase(s string) string {
    result := make([]rune, len(s))

    for i, char := range s {
        if char >= 'a' && char <= 'z' {
            result[i] = char - 'a' + 'A'
        } else {
            result[i] = char
        }
    }

    return string(result)
}
