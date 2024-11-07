package sprint
func StrSplitBy(s, sep string) []string {
    if s == "" {
        return []string{}
    }

    var substrings []string
    var substring string

    // Iterate over each character in the string s
    for i := 0; i < len(s); i++ {
        // Check if the current position in the string matches the separator sep
        if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
            // If it does, add the accumulated substring to the slice of substrings
            if substring != "" {
                substrings = append(substrings, substring)
            }
            // Reset the substring variable
            substring = ""
            // Move the index to the end of the separator sep
            i += len(sep) - 1
        } else {
            // If not, add the current character to the accumulated substring
            substring += string(s[i])
        }
    }

    // Add the last accumulated substring to the slice of substrings
    if substring != "" {
        substrings = append(substrings, substring)
    }

    return substrings
}
