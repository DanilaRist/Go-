package sprint
func SubstrIndex(s string, toFind string) int {
    // Iterate over each character in the string s
    for i := 0; i < len(s)-len(toFind)+1; i++ {
        found := true
        // Check if the substring toFind is present starting from index i in s
        for j := 0; j < len(toFind); j++ {
            if s[i+j] != toFind[j] {
                found = false
                break
            }
        }
        // If the substring toFind is found, return its index
        if found {
            return i
        }
    }
    // If toFind is not found in s, return -1
    return -1
}
