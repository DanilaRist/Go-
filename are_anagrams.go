package sprint

import (
    "sort"
    "strings"
)

func AreAnagrams(str1, str2 string) bool {
    // Remove whitespaces and convert strings to lowercase
    str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
    str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

    // Check if lengths of strings are different, if yes, they cannot be anagrams
    if len(str1) != len(str2) {
        return false
    }

    // Convert strings to slices of runes for sorting
    runes1 := []rune(str1)
    runes2 := []rune(str2)

    // Sort the rune slices
    sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
    sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

    // Check if the sorted rune slices are equal
    for i := range runes1 {
        if runes1[i] != runes2[i] {
            return false
        }
    }

    return true
}