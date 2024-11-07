package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
    filteredArr := make([][]int, 0)

    // Iterate over each subarray
    for _, subarr := range arr {
        sum := 0
        // Calculate the sum of elements in the subarray
        for _, num := range subarr {
            sum += num
        }
        // If the sum is greater than or equal to the limit, include the subarray in the result
        if sum >= limit {
            filteredArr = append(filteredArr, subarr)
        }
    }

    return filteredArr
}
