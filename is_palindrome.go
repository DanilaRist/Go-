package sprint

func IsPalindrome(s string) bool {
    // Function to check if a character is a white space
    isWhiteSpace := func(c byte) bool {
        return c == ' '
    }

    // Remove white spaces from the string
    var cleaned string
    for i := 0; i < len(s); i++ {
        if !isWhiteSpace(s[i]) {
            cleaned += string(s[i])
        }
    }

    // Initialize pointers for the start and end of the string
    i, j := 0, len(cleaned)-1

    // Iterate through the cleaned string from both ends
    for i < j {
        // If characters at pointers don't match, return false
        if cleaned[i] != cleaned[j] {
            return false
        }
        // Move pointers inward
        i++
        j--
    }

    // If the loop completes without returning false, the string is a palindrome
    return true
}