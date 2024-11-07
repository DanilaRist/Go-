package sprint
func IsPrime(n int) bool {
    if n <= 1 {
        return false // 1 is not a prime number
    }
    if n == 2 {
        return true // 2 is a prime number
    }
    if n%2 == 0 {
        return false // Even numbers greater than 2 are not prime
    }
    
    // Check divisibility only up to square root
    for i := 3; i*i <= n; i += 2 {
        if n%i == 0 {
            return false // Found a divisor, not a prime number
        }
    }
    
    return true // No divisors found, prime number
}
