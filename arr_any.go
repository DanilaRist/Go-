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

// ArrAny returns true for a string slice if, when that string slice is passed through an f function, at least one element returns true.
func ArrAny(f func(string) bool, a []string) bool {
    for _, s := range a {
        if f(s) {
            return true
        }
    }
    return false
}
