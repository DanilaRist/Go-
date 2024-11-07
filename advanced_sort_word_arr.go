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

// AdvancedSortWordArr sorts a slice of strings based on the provided comparison function f.
func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
    for i := 0; i < len(a)-1; i++ {
        for j := 0; j < len(a)-i-1; j++ {
            if f(a[j], a[j+1]) == 1 {
                a[j], a[j+1] = a[j+1], a[j]
            }
        }
    }
    return a
}
