package sprint

// ConvertIntToString converts an integer to a string manually
func ConvertIntToString(n int) string {
    // Maximum number of digits for a 32-bit integer is 10
    // (including the sign for negative numbers)
    const maxDigits = 10
    buf := make([]byte, 0, maxDigits+1) // +1 for potential negative sign

    // Handle negative numbers
    if n < 0 {
        buf = append(buf, '-')
        n = -n
    }

    // Convert each digit to a character and prepend to the buffer
    if n == 0 {
        buf = append(buf, '0')
    } else {
        for n > 0 {
            digit := byte('0' + n%10)
            buf = append(buf, digit)
            n /= 10
        }
        // Reverse the buffer to get the correct order of digits
        for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
            buf[i], buf[j] = buf[j], buf[i]
        }
    }

    return string(buf)
}

func GenerateRange(min, max int) []int {
    if min >= max {
        return nil
    }

    var result []int
    for i := min; i < max; i++ {
        result = append(result, i)
    }

    return result
}

func IntSliceToString(slice []int) string {
    if slice == nil {
        return "[]"
    }

    var result string
    result += "["
    for i, v := range slice {
        if i > 0 {
            result += " "
        }
        result += ConvertIntToString(v)
    }
    result += "]"
    return result
}

func main() {
    result := GenerateRange(5, 10)
    println(IntSliceToString(result)) // Output: [5 6 7 8 9]

    result = GenerateRange(10, 5)
    println(IntSliceToString(result)) // Output: []
}
