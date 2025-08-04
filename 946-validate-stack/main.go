package validatestack

func validateStackSequences(pushed []int, popped []int) bool {
	// Create a stack to simulate operations
	stack := make([]int, 0, len(pushed))
	popIndex := 0 // Index to track position in popped array

	// Process each value in pushed array
	for _, val := range pushed {
		// Push the current value onto stack
		stack = append(stack, val)

		// Try to pop values that match the popped sequence
		for len(stack) > 0 && popIndex < len(popped) && stack[len(stack)-1] == popped[popIndex] {
			// Pop from stack
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}

	// Check if we were able to pop all elements in the correct sequence
	return len(stack) == 0
}
