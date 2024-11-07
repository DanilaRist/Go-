package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	// Ensure from is smaller than to
	if from > to {
		from, to = to, from
	}

	// Ensure from and to are within bounds
	if from < 0 {
		from = 0
	} else if from >= len(arr) {
		from = len(arr) - 1
	}
	if to < 0 {
		to = 0
	} else if to >= len(arr) {
		to = len(arr) - 1
	}

	// Remove elements between from and to, inclusive
	result := make([]float64, 0, len(arr)-(to-from))
	result = append(result, arr[:from]...)
	if to+1 < len(arr) {
		result = append(result, arr[to:]...)
	}

	return result
}