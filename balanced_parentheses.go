package sprint
func BalancedParentheses(input string) bool {
	stack := make([]rune, 0)

	// Define a mapping of opening to closing parentheses
	parenthesesMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// Iterate through each character in the input string
	for _, char := range input {
		// If the character is an opening parenthesis, push it onto the stack
		if isOpeningParenthesis(char) {
			stack = append(stack, char)
		} else if isClosingParenthesis(char) { // If the character is a closing parenthesis
			// If the stack is empty or the closing parenthesis does not match the top of the stack, return false
			if len(stack) == 0 || parenthesesMap[stack[len(stack)-1]] != char {
				return false
			}
			// Pop the corresponding opening parenthesis from the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is empty, all parentheses are balanced, otherwise, return false
	return len(stack) == 0
}

func isOpeningParenthesis(char rune) bool {
	return char == '(' || char == '[' || char == '{'
}

func isClosingParenthesis(char rune) bool {
	return char == ')' || char == ']' || char == '}'
}
