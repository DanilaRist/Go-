package sprint
func PascalsTriangle(n int) [][]int {
    // Initialize the Pascal's triangle with the first level
    triangle := [][]int{{1}}

    // Generate subsequent levels of Pascal's triangle
    for i := 1; i < n; i++ {
        // Initialize the current level with the leftmost value
        level := []int{1}

        // Compute the values of the current level based on the previous level
        for j := 1; j < i; j++ {
            // Each value is the sum of the values above and to the left
            value := triangle[i-1][j-1] + triangle[i-1][j]
            level = append(level, value)
        }

        // Add the rightmost value to the current level
        level = append(level, 1)

        // Add the current level to the Pascal's triangle
        triangle = append(triangle, level)
    }

    return triangle
}