package sprint
import (
    "strconv"
    "strings"
)

var solutions []string

func EightQueensSolver() string {
    solutions = []string{} // Initialize solutions slice
    board := make([]int, 8) // Initialize the board with 8 columns
    placeQueen(board, 0) // Start placing queens recursively from the first row
    return strings.Join(solutions, "\n") // Join solutions into a single string separated by newlines
}

// Recursive function to place queens on the board
func placeQueen(board []int, row int) {
    if row == 8 { // If we've placed 8 queens, record the solution
        solution := ""
        for _, col := range board {
            solution += strconv.Itoa(col + 1) // Convert column index to 1-based index
        }
        solutions = append(solutions, solution)
        return
    }
    for col := 0; col < 8; col++ { // Try placing a queen in each column of the current row
        if isValid(board, row, col) { // Check if it's safe to place a queen at this position
            board[row] = col
            placeQueen(board, row+1) // Recur to the next row
        }
    }
}

// Function to check if it's safe to place a queen at a given position
func isValid(board []int, row, col int) bool {
    for i := 0; i < row; i++ { // Check previous rows
        if board[i] == col || // Same column
            board[i]+i == col+row || // Same major diagonal
            board[i]-i == col-row { // Same minor diagonal
            return false
        }
    }
    return true
}

func main() {
    // Example usage
    solutions := EightQueensSolver()
    println(solutions)
}