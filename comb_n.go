package sprint


func CombN(n int) []string {
	var result []string
	backtrack("", 0, n, &result)
	return result
}

func backtrack(combination string, start, n int, result *[]string) {
	if len(combination) == n {
		*result = append(*result, combination)
		return
	}
	for i := start; i <= 9; i++ {
		backtrack(combination+string('0'+i), i+1, n, result)
	}
}

func main() {
	// Example usage
	combinations := CombN(3)
	for _, comb := range combinations {
		println(comb)
	}
}
