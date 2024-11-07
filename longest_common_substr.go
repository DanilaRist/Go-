package sprint

func LongestCommonSubstr(str1, str2 string) string {
	var longest string

	// Create a 2D slice to store the lengths of common substrings
	lengths := make([][]int, len(str1)+1)
	for i := range lengths {
		lengths[i] = make([]int, len(str2)+1)
	}

	// Iterate over the strings and fill the lengths slice
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				lengths[i][j] = lengths[i-1][j-1] + 1
				if lengths[i][j] > len(longest) {
					longest = str1[i-lengths[i][j]:i]
				}
			}
		}
	}

	return longest
}
