package sprint
func SortWordArr(a []string) []string {
    // Bubble sort algorithm
    for i := 0; i < len(a)-1; i++ {
        for j := 0; j < len(a)-i-1; j++ {
            // Compare strings based on ASCII values
            if a[j] > a[j+1] {
                // Swap elements if they are in the wrong order
                a[j], a[j+1] = a[j+1], a[j]
            }
        }
    }
    return a
}