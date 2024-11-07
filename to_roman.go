package sprint
func ToRoman(num int) string {
    // Input validation
    if num < 1 || num > 3999 {
        return "Invalid input"
    }

    // Define Roman numeral symbols and their corresponding values
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

    // Initialize the result string
    var result string

    // Iterate through the values and symbols to construct the Roman numeral
    for i := 0; i < len(values); i++ {
        for num >= values[i] {
            result += symbols[i]
            num -= values[i]
        }
    }

    return result
}