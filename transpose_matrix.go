package sprint
func TransposeMatrix(matrix [][]int) [][]int {
    // Get the dimensions of the original matrix
    numRows := len(matrix)
    if numRows == 0 {
        return [][]int{}
    }
    numCols := len(matrix[0])

    // Create a new matrix with swapped dimensions
    transposed := make([][]int, numCols)
    for i := range transposed {
        transposed[i] = make([]int, numRows)
    }

    // Fill the transposed matrix with elements from the original matrix
    for i := 0; i < numRows; i++ {
        for j := 0; j < numCols; j++ {
            transposed[j][i] = matrix[i][j]
        }
    }

    return transposed
}
