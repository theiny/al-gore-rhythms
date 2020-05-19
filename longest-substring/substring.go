package substring

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var new string
	var longest int

	for i := 0; i < len(s); i++ {
		firstLetter := rune(s[i])
		if !strings.ContainsRune(new, firstLetter) {
			new += string(firstLetter)
		}
	Inner:
		for j := i + 1; j < len(s); j++ {
			lastLetter := rune(s[j])
			if !strings.ContainsRune(new, lastLetter) {
				new += string(lastLetter)
			} else {
				if len(new) > longest {
					longest = len(new)
				}
				new = ""
				break Inner
			}
		}
	}

	if longest > len(new) {
		return longest
	}

	return len(new)
}
