package parantheses

func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	var stack []rune

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, m[c])
		} else if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] != c {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
