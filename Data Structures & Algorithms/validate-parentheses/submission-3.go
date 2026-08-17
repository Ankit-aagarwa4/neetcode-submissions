func isValid(s string) bool {
	// convert the string to slice of characters
	chars := []rune(s)
	stack := make([]rune, len(chars))
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	i := -1
	for _, val := range chars {
		if val == '(' || val == '[' || val == '{' {
			i++
			stack[i] = val
		} else {
			// check if val as key in brackets and if it matches the top of the stack
			if _, ok := brackets[val]; ok {
				if i == -1 {
					return false
				}
				if stack[i] != brackets[val] {
					return false
				}
				i--
			}
		}
	}

	return i == -1
}
