package sprint

import (
    "strings"
)

func ConvertAnyToDec(s, base string) int {
    // Validate the base
    if len(base) < 2 || strings.ContainsAny(base, "+-") {
        return 0
    }

    // Create a map to store the index of each character in the base string
    indexMap := make(map[rune]int)
    for i, ch := range base {
        indexMap[ch] = i
    }

    // Initialize the result to 0
    result := 0

    // Iterate over each digit in the numeric string and calculate its value in decimal
    for _, digit := range s {
        index, ok := indexMap[digit]
        if !ok {
            return 0 // The digit is not in the base, return 0
        }
        result = result*len(base) + index
    }

    return result
}

func ConvertDecToAny(n int, base string) string {
    // Validate the base
    if len(base) < 2 || strings.ContainsAny(base, "+-") {
        return "NV"
    }

    // Convert the decimal number to the desired base
    result := ""
    baseLength := len(base)
    for n > 0 {
        digit := n % baseLength
        result = string(base[digit]) + result
        n /= baseLength
    }

    // If the result is empty (n was 0), return "0"
    if result == "" {
        result = "0"
    }

    return result
}

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
    // Validate the bases
    if len(baseFrom) < 2 || len(baseTo) < 2 || strings.ContainsAny(baseFrom, "+-") || strings.ContainsAny(baseTo, "+-") {
        return "NV"
    }

    // Convert nbr from baseFrom to decimal
    decimal := ConvertAnyToDec(nbr, baseFrom)

    // If the conversion failed (returned 0), return "NV"
    if decimal == 0 {
        return "NV"
    }

    // Convert decimal to baseTo
    return ConvertDecToAny(decimal, baseTo)
}
