package sprint


// AlphaNumber converts an integer to a string where digits are replaced by lowercase letters.
func AlphaNumber(n int) string {
    // Handle zero separately
    if n == 0 {
        return "a"
    }

    // Convert the integer to a string
    str := ""

    // Handle negative numbers
    if n < 0 {
        // Convert the absolute value of the negative number
        absN := ""
        for n != 0 {
            digit := -(n % 10) // Convert the digit to a negative value
            letter := 'a' + digit
            absN = string(letter) + absN // Concatenate the letter at the beginning
            n /= 10
        }

        // Add the minus sign at the beginning
        str = "-" + absN
    } else {
        // Convert each digit to its corresponding letter
        for n != 0 {
            digit := n % 10
            letter := 'a' + digit
            str = string(letter) + str // Concatenate at the beginning
            n /= 10
        }
    }

    return str
}

func main() {
    // Test cases
    println(AlphaNumber(-1280))
}
