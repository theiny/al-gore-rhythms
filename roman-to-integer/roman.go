package roman

func romanToInt(s string) int {
	var count int
	var prev string
	m := map[string]int{}
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	for i := len(s) - 1; i >= 0; i-- {
		letter := string(s[i])

		switch {
		case letter == "I" && (prev == "V" || prev == "X"):
			count -= m[letter]
		case letter == "X" && (prev == "L" || prev == "C"):
			count -= m[letter]
		case letter == "C" && (prev == "D" || prev == "M"):
			count -= m[letter]
		default:
			count += m[letter]
		}

		prev = letter
	}

	return count
}
