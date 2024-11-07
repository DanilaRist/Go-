package sprint
func Payout(amount int, denominations []int) (payout []int) {
    // Perform bubble sort in descending order (simple implementation)
    for i := 0; i < len(denominations)-1; i++ {
        for j := 0; j < len(denominations)-i-1; j++ {
            if denominations[j] < denominations[j+1] {
                // Swap elements if they are in the wrong order
                denominations[j], denominations[j+1] = denominations[j+1], denominations[j]
            }
        }
    }

    // Initialize result array
    payout = []int{}

    // Iterate through denominations
    for _, d := range denominations {
        // Repeat adding the denomination until it exceeds the amount
        for amount >= d {
            // Add the denomination to the result
            payout = append(payout, d)
            // Deduct the denomination from the amount
            amount -= d
        }
    }

    // If amount is not 0, it means payout cannot be made, return empty array
    if amount != 0 {
        payout = []int{}
    }

    return payout
}