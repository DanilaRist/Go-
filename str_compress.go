package sprint
import "fmt"

func StrCompress(input string) string {
    // Edge case: empty input or single character
    if len(input) <= 1 {
        return input
    }

    // Initialize variables
    var compressed string
    currentChar := input[0]
    count := 1

    // Iterate over the characters of the input string starting from the second character
    for i := 1; i < len(input); i++ {
        // If the current character is the same as the previous character, increase the count
        if input[i] == currentChar {
            count++
        } else {
            // If the current character is different from the previous character,
            // append the count and the character to the compressed string
            if count > 1 {
                compressed += fmt.Sprintf("%d%s", count, string(currentChar))
            } else {
                compressed += string(currentChar)
            }
            // Reset count and update currentChar to the new character
            count = 1
            currentChar = input[i]
        }
    }

    // Append the last count and character to the compressed string
    if count > 1 {
        compressed += fmt.Sprintf("%d%s", count, string(currentChar))
    } else {
        compressed += string(currentChar)
    }

    return compressed
}