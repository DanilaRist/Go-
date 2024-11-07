package sprint
// IsNegative returns true if the given integer is negative, and false otherwise.
func IsNegative(n int) bool {
    return n < 0
}

// IsPrime returns true if the given integer is a prime number, and false otherwise.
func IsPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n <= 3 {
        return true
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }
    for i := 5; i*i <= n; i += 6 {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
    }
    return true
}

// ArrMap applies a function of type func(int) bool on each element of an int slice and returns a slice of all the return values.
func ArrMap(f func(int) bool, a []int) []bool {
    result := make([]bool, len(a))
    for i, v := range a {
        result[i] = f(v)
    }
    return result
}