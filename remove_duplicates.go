package sprint
func RemoveDuplicates(arr []int) []int {
    // Create a map to store unique elements
    uniqueMap := make(map[int]bool)
    result := []int{}

    // Iterate through the array
    for _, num := range arr {
        // If the element is not in the map, add it to the result slice
        if !uniqueMap[num] {
            uniqueMap[num] = true
            result = append(result, num)
        }
    }

    return result
}
