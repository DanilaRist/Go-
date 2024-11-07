package sprint
// StrCompare compares two strings lexicographically.
func StrCompare(a, b string) int {
    if a < b {
        return -1
    }
    if a > b {
        return 1
    }
    return 0
}

// IsSorted returns true if the slice of string is sorted according to a provided comparison function, and false otherwise.
func IsSorted(f func(a, b string) int, arr []string) bool {
    n := len(arr)
    ascending := true
    descending := true
    equal := true

    for i := 1; i < n; i++ {
        cmp := f(arr[i-1], arr[i])

        if cmp > 0 {
            ascending = false
        } else if cmp < 0 {
            descending = false
        }

        if cmp != 0 {
            equal = false
        }
    }

    return ascending || descending || equal
}

