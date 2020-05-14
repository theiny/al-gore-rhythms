package palindrome

func isPalindrome(x int) bool {
	var new int
	original := x

	for x > 0 {
		remainder := x % 10 // remainder = 1
		new *= 10           // new = 0
		new += remainder    // new = 1
		x /= 10             // x = 12
	}

	if original == new {
		return true
	}

	return false
}
