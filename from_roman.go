package sprint
func FromRoman(s string) int {
    // Create a map to store the values of Roman numerals
    romanMap := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    // Initialize the result
    result := 0

    // Iterate through the Roman numeral string
    for i := 0; i < len(s); i++ {
        // Get the value of the current Roman numeral
        value := romanMap[s[i]]

        // If the current numeral is smaller than the next one, subtract its value from the result
        if i+1 < len(s) && romanMap[s[i]] < romanMap[s[i+1]] {
            result -= value
        } else {
            // Otherwise, add its value to the result
            result += value
        }
    }

    return result
}