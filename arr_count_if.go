package sprint
// IsLower returns true if the string given contains only lowercase characters, and false otherwise.
func IsLower(s string) bool {
    for _, r := range s {
        if !(r >= 'a' && r <= 'z') {
            return false
        }
    }
    return true
}

// IsNumeric returns true if the string given contains only numeric characters, and false otherwise.
func IsNumeric(s string) bool {
    for _, r := range s {
        if !(r >= '0' && r <= '9') {
            return false
        }
    }
    return true
}

// IsUpper returns true if the string given contains only uppercase characters, and false otherwise.
func IsUpper(s string) bool {
    for _, r := range s {
        if !(r >= 'A' && r <= 'Z') {
            return false
        }
    }
    return true
}

// IsAlphanumeric returns true if the string contains only alphanumeric characters, and false otherwise.
func IsAlphanumeric(s string) bool {
    for _, r := range s {
        if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9') {
            return false
        }
    }
    return true
}

// ArrCountIf returns the number of elements in a string slice for which the f function returns true.
func ArrCountIf(f func(string) bool, tab []string) int {
    count := 0
    for _, s := range tab {
        if f(s) {
            count++
        }
    }
    return count
}