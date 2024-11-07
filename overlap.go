package sprint

func Overlap(arr1, arr2 []int) []int {
	// Check if either input array is empty
	if len(arr1) == 0 || len(arr2) == 0 {
		return []int{}
	}

	// Special case: if both input arrays have only one element
	if len(arr1) == 1 && len(arr2) == 1 {
		if arr1[0] == arr2[0] {
			return []int{arr1[0]}
		}
		return []int{}
	}

	var result []int
	freqMap := make(map[int]int)

	for _, num := range arr1 {
		freqMap[num]++
	}

	for _, num := range arr2 {
		if freqMap[num] > 0 {
			result = append(result, num)
			freqMap[num]--
		}
	}

	bubbleSort(result)
	return result
}

func bubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
