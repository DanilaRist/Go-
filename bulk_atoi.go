package sprint

func BulkAtoi(arr []string) []int {
    intArr := make([]int, len(arr))
    for i, str := range arr {
        intArr[i] = StrToInt(str)
    }
    return intArr
}

// StrToInt function copied here
func StrToInt(str string) int {
    result := 0
    sign := 1
    for i, char := range str {
        if i == 0 && char == '-' {
            sign = -1
            continue
        }
        digit := int(char - '0')
        if digit < 0 || digit > 9 {
            // Invalid character, return 0
            return 0
        }
        result = result*10 + digit
    }
    return sign * result
}