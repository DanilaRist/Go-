package sprint 
func BalanceOut(arr []bool) []bool {
    trueCount, falseCount := 0, 0
    
    // Count the number of true and false values in the array
    for _, value := range arr {
        if value {
            trueCount++
        } else {
            falseCount++
        }
    }
    
    // Calculate the difference between the counts
    diff := abs(trueCount - falseCount)
    
    // Create a new slice to hold the balanced array
    balancedArr := make([]bool, len(arr))
    copy(balancedArr, arr)
    
    // Add the necessary number of booleans to balance the array
    for i := 0; i < diff; i++ {
        if trueCount > falseCount {
            balancedArr = append(balancedArr, false)
            falseCount++
        } else {
            balancedArr = append(balancedArr, true)
            trueCount++
        }
    }
    
    return balancedArr
}

// Helper function to find the absolute value
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
