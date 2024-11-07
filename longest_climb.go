package sprint
func LongestClimb(arr []int) []int {
    // Edge case: if the array is empty, return an empty slice
    if len(arr) == 0 {
        return []int{}
    }

    // Initialize variables to track the longest subarray and the current subarray
    var longestSubarray, currentSubarray []int

    // Iterate through the array
    for i := 0; i < len(arr)-1; i++ {
        // If the current element is less than or equal to the next element,
        // add it to the current subarray
        if arr[i] <= arr[i+1] {
            // If the current subarray is empty, add the current element
            if len(currentSubarray) == 0 {
                currentSubarray = append(currentSubarray, arr[i])
            }
            currentSubarray = append(currentSubarray, arr[i+1])
        } else {
            // If the current element is not less than or equal to the next element,
            // update the longest subarray if the current subarray is longer
            if len(currentSubarray) > len(longestSubarray) {
                longestSubarray = make([]int, len(currentSubarray))
                copy(longestSubarray, currentSubarray)
            }
            // Reset the current subarray
            currentSubarray = []int{}
        }
    }

    // Check if the last subarray is longer than the longest subarray
    if len(currentSubarray) > len(longestSubarray) {
        longestSubarray = make([]int, len(currentSubarray))
        copy(longestSubarray, currentSubarray)
    }

    return longestSubarray
}