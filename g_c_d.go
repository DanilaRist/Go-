package sprint
func GCD(a, b int) int {
    // Keep finding remainder until b becomes 0
    for b != 0 {
        a, b = b, a%b
    }
    // GCD will be stored in a
    return a
}
