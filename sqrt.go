package sprint

func Sqrt(n int) int {
    if n < 0 {
        return 0 // Square root of negative numbers is not defined
    }

    for i := 1; i*i <= n; i++ {
        if i*i == n {
            return i // Found the integer square root
        }
    }
    
    return 0 // No integer square root found
}