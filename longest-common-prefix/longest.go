package longest

func longestCommonPrefix(strs []string) string {
	var length int
	// We want to find the length of the shortest word, since we don't need to check past this.
	for i, str := range strs {
		if i == 0 {
			length = len(str)
			continue
		}

		if len(str) < length {
			length = len(str)
		}
	}

	var prefix string
	var prev string
	var letter string
	var match bool

Outer:
	for i := 0; i < length; i++ {
		for j, str := range strs {
			letter = string(str[i])

			if j == 0 {
				prev = letter
			}

			if letter == prev {
				match = true
			} else {
				match = false
				break Outer
			}

			prev = letter

		}
		if match {
			prefix += letter
		}

	}

	return prefix
}
