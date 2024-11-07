package sprint
func ToCapitalCase(s string) string {
    // Convert the string to a rune slice
    runes := []rune(s)
    
    // Initialize a flag to indicate if the next character should be capitalized
    capitalizeNext := true
    
    // Iterate over each character in the string
    for i, char := range runes {
        // Check if the character is a letter or digit
        if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
            // If the next character should be capitalized
            if capitalizeNext {
                // Capitalize the character
                if 'a' <= char && char <= 'z' {
                    runes[i] = char - 'a' + 'A'
                }
                // The next character should not be capitalized unless it is alphanumeric
                capitalizeNext = false
            } else {
                // Convert the character to lowercase
                if 'A' <= char && char <= 'Z' {
                    runes[i] = char - 'A' + 'a'
                }
            }
        } else {
            // If the character is not a letter or digit, reset the flag to capitalize the next character
            capitalizeNext = true
        }
    }
    
    // Convert the modified rune slice back to a string
    return string(runes)
}
