package sprint
func GCD(a, b int) int {
    // Calculate GCD using Euclidean algorithm
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func LCM(a, b int) int {
    // Calculate LCM using the formula: LCM(a, b) = (a * b) / GCD(a, b)
    return (a * b) / GCD(a, b)
}
